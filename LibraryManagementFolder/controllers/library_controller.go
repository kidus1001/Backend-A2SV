package controllers

import (
	"fmt"
	"librarymanagementfolder/models"
	"librarymanagementfolder/services"
)

func StartLibraryApp() {
	library := services.NewLibrary()

	for {
		fmt.Println("*** Welcome to the library management system! ***")
		fmt.Println("\n1. To add a new book")
		fmt.Println("2. To remove an existing book")
		fmt.Println("3. To borrow a book")
		fmt.Println("4. To return a book")
		fmt.Println("5. To list all the available books")
		fmt.Println("6. To list all borrowed books by a member")
		fmt.Print("Enter a number: ")
		var numChosen string
		fmt.Scanln(&numChosen)

		switch numChosen {
		case "1":

			fmt.Println("Enter the ID of the book: ")
			var ID int
			fmt.Scanln(&ID)

			fmt.Println("Enter the title of the book: ")
			var title string
			fmt.Scanln(&title)

			fmt.Println("Enter the author of the book: ")
			var author string
			fmt.Scanln(&author)

			var bookToAdd models.Book
			bookToAdd.ID = ID
			bookToAdd.Title = title
			bookToAdd.Author = author
			bookToAdd.Status = "Available"

			library.AddBook(bookToAdd)
		case "2":
			fmt.Println("Enter the ID of the book to remove: ")
			var ID int
			fmt.Scanln(&ID)
			library.RemoveBook(ID)

		case "3":
			fmt.Println("Enter the ID of the book: ")
			var bookID int
			fmt.Scanln(&bookID)

			fmt.Println("Enter the ID of the member: ")
			var ID int
			fmt.Scanln(&ID)
			library.BorrowBook(bookID, ID)

		case "4":
			fmt.Println("Enter the ID of the book: ")
			var bookID int
			fmt.Scanln(&bookID)

			fmt.Println("Enter the ID of the member: ")
			var ID int
			fmt.Scanln(&ID)
			library.ReturnBook(bookID, ID)

		case "5":
			availableBooks := library.ListAvailableBooks()

			for _, i := range availableBooks {
				fmt.Println(i.Title)
			}

		case "6":
			fmt.Println("Enter the ID of the member: ")
			var ID int
			fmt.Scanln(&ID)
			library.ListBorrowedBooks(ID)
		default:
			fmt.Println("Invalid option!")
		}
	}
}
