# Library Management System

## Overview
This is a console-based application written in Go to manage library operations. It demonstrates the use of structs, interfaces, maps, slices, and modular file structure.

## Features
1. **Add Book**: Add a new book to the library inventory.
2. **Remove Book**: Permanently remove a book from the system.
3. **Borrow Book**: Assign a book to a member (changes status to 'Borrowed').
4. **Return Book**: Return a book to the library (changes status to 'Available').
5. **List Available**: Show all books currently in the library.
6. **List Borrowed**: Show all books currently held by a specific member.

## How to Run
1. Ensure Go is installed.
2. Navigate to the `library_management` folder.
3. Run the following command:
   ```bash
   go run main.go