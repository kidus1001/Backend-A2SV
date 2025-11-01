package services

import (
	"errors"
	"fmt"
	"librarymanagementfolder/models"
)

type Library struct {
	allBooks   map[int]models.Book
	allMembers map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		allBooks:   make(map[int]models.Book),
		allMembers: make(map[int]models.Member),
	}
}

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

func (l *Library) AddBook(book models.Book) {
	bookId := book.ID
	if _, exists := l.allBooks[bookId]; !exists {
		l.allBooks[bookId] = book
		fmt.Println("Added: ", l.allBooks[book.ID].Title)
	} else {
		fmt.Println("Book already exists!")
	}
}

func (l *Library) RemoveBook(ID int) {
	delete(l.allBooks, ID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	if l.allBooks[bookID].Status == "Available" {
		l.allMembers[memberID].BorrowedBooks[bookID] = l.allBooks[bookID]
		fmt.Println("Book borrowed successfully!")
		l.allMembers[memberID].BorrowedBooks[bookID].Status = "Borrowed"
		return nil
	} else {
		return errors.New("book is not available currently")
	}
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	member, memberExists := l.allMembers[memberID]

	if !memberExists {
		return errors.New("member with ID not found")
	}

	book := member.BorrowedBooks[bookID]
	if book.Status == "Available" {
		return errors.New("book with the given ID is not borrowed by this member")
	}
	book.Status = "Available"
	l.allBooks[bookID] = book

	updatedBorrowedBooks := []models.Book{}
	for _, n := range l.allMembers[memberID].BorrowedBooks {
		if n != l.allBooks[bookID] {
			updatedBorrowedBooks = append(updatedBorrowedBooks, n)
		}
	}

	member.BorrowedBooks = updatedBorrowedBooks
	l.allMembers[memberID] = member

	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	available := []models.Book{}

	for _, n := range l.allBooks {
		if n.Status == "Available" {
			available = append(available, n)
		}
	}

	return available
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	return l.allMembers[memberID].BorrowedBooks
}
