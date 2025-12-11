package services

import (
	"errors"
	"fmt"
	"task3/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	AddMember(member models.Member)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	if _, exists := l.books[book.ID]; exists {
		fmt.Printf("Book with ID %d already exists.\n", book.ID)
		return
	}
	l.books[book.ID] = book
	fmt.Println("Book added successfully.")
}

func (l *Library) AddMember(member models.Member) {
	if _, exists := l.members[member.ID]; exists {
		fmt.Printf("Member with ID %d already exists.\n", member.ID)
		return
	}
	l.members[member.ID] = member
	fmt.Println("Member registerd successfully.")
}

func (l *Library) RemoveBook(bookID int) {
	if _, exists := l.books[bookID]; exists {
		fmt.Println("Book not found.")
		return
	}
	delete(l.books, bookID)
	fmt.Println("Book removed successfully")
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, bookExixts := l.books[bookID]
	member, memberExists := l.members[memberID]

	if !bookExixts {
		return errors.New("book not found")
	}
	if !memberExists {
		return errors.New("member not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	// Update book status
	book.Status = "Borrowed"
	l.books[bookID] = book // Update map

	// Update memeber
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.members[memberID] = member // Update map

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, bookExists := l.books[bookID]
	member, memberExists := l.members[memberID]

	if !bookExists {
		return errors.New("book not found in library records")
	}
	if !memberExists {
		return errors.New("member not found")
	}

	// Find book in member's borrowed list
	foundIndex := -1
	for index, book := range member.BorrowedBooks {
		if book.ID == bookID {
			foundIndex = index
			break
		}
	}

	if foundIndex == -1 {
		return errors.New("this member does not have this book in borrowed")
	}

	// Remove from member's list
	member.BorrowedBooks = append(member.BorrowedBooks[:foundIndex], member.BorrowedBooks[foundIndex+1:]...)
	l.members[memberID] = member

	// Update book status
	book.Status = "Available"
	l.books[bookID] = book

	return nil
}

func (l* Library) ListAvailableBooks() []models.Book {
	var available []models.Book
	for _, book := range l.books {
		if book.Status == "Available" {
			available = append(available, book)
		}
	}
	return available
}

func (l* Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.members[memberID]
	if !exists {
		fmt.Println("Member not found")	
		return nil
	}
	return member.BorrowedBooks
}
