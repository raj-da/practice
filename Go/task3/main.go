package main

import (
	"fmt"
	"task3/controllers"
	"task3/models"
	"task3/services"
)

func main() {
	libService := services.NewLibrary()

	// 2. Pre-load some dummy data for testing
	libService.AddBook(models.Book{ID: 1, Title: "The Go Programming Language", Author: "Alan Donovan", Status: "Available"})
	libService.AddBook(models.Book{ID: 2, Title: "Clean Code", Author: "Robert C. Martin", Status: "Available"})
	
	libService.AddMember(models.Member{ID: 101, Name: "Alice Smith", BorrowedBooks: []models.Book{}})
	libService.AddMember(models.Member{ID: 102, Name: "Bob Jones", BorrowedBooks: []models.Book{}})

	fmt.Println("System initialized with Members: 101 (Alice), 102 (Bob)")

	libContoller := controllers.NewLibraryController(libService)
	libContoller.Run()
}