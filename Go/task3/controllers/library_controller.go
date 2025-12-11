package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task3/models"
	"task3/services"
)

type LibraryController struct {
	service services.LibraryManager
	reader  *bufio.Reader
}

func NewLibraryController(service services.LibraryManager) *LibraryController {
	return &LibraryController{
		service: service,
		reader:  bufio.NewReader(os.Stdin),
	}
}

func (lc *LibraryController) readInt(prompt string) int {
	fmt.Print(prompt)
	input, _ := lc.reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return -1
	}
	return num
}

func (lc *LibraryController) readString(prompt string) string {
	fmt.Print(prompt)
	input, _ := lc.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (lc *LibraryController) Run() {
	for {
		fmt.Println("\n--- Library Management System ---")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Exit")
		choice := lc.readInt("Enter your choice: ")

		switch choice {
		case 1:
			id := lc.readInt("Enter Book ID: ")
			title := lc.readString("Enter Title: ")
			author := lc.readString("Enter Author: ")
			newBook := models.Book{ID: id, Title: title, Author: author, Status: "Available"}
			lc.service.AddBook(newBook)

		case 2:
			id := lc.readInt("Enter Book ID to remove: ")
			lc.service.RemoveBook(id)

		case 3:
			bookID := lc.readInt("Enter Book ID: ")
			memberID := lc.readInt("Enter Member ID: ")
			err := lc.service.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Book borrowed successfully.")
			}

		case 4:
			bookID := lc.readInt("Enter Book ID: ")
			memberID := lc.readInt("Enter Member ID: ")
			err := lc.service.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Book returned successfully.")
			}

		case 5:
			books := lc.service.ListAvailableBooks()
			fmt.Println("\n--- Available Books ---")
			if len(books) == 0 {
				fmt.Println("No books available.")
			}
			for _, b := range books {
				fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
			}

		case 6:
			memberID := lc.readInt("Enter Member ID: ")
			books := lc.service.ListBorrowedBooks(memberID)

			if len(books) == 0 {
				fmt.Println("No books borrowed by this member.")
			}
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid Input!")
		}
	}
}
