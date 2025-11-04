package main

import (
	"fmt"
	"errors"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Status string //Available or borrowed
}

type Member struct {
	ID            int
	Name          string
	BorrowedBooks []Book
}

type LibraryManager interface {
	AddBook(book Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []Book
	ListBorrowedBooks(memberID int) []Book
}

type Library struct {
	allBooks   map[int]Book
	allMembers map[int]Member
}

func (l *Library) AddBook(book Book) {
	bookId := book.ID
	if _, exists := l.allBooks[bookId]; !exists {
		l.allBooks[bookId] = book
		fmt.Println("Added: ", l.allBooks[book.ID].Title)
	} else {
		fmt.Println("Book already exists!")
	}
}

func (l *Library) RemoveBook (ID int) {
	delete(l.allBooks, ID)
}

func (l *Library) BorrowBook (bookID int, memberID int) error {
	if (l.allBooks[bookID].Status == "Available") {
		l.allMembers[memberID].BorrowedBooks[bookID] = l.allBooks[bookID] 
		fmt.Println("Book borrowed successfully!")
		return nil
	} else {
		return errors.New("Book is not available currently!")
	}	
}

func (l *Library) ReturnBook (bookID int, memberID int) error {
	member, memberExists := l.allMembers[memberID]

	if !memberExists {
		return errors.New ("Member with ID not found!")
	}

	book := member.BorrowedBooks[bookID]
	if book.Status ==  "Available" {
		return errors.New("Book with the given ID is not borrowed by this member")
	}
	book.Status = "Available"
	l.allBooks[bookID] = book

	updatedBorrowedBooks := []Book {}
	for _, n := range l.allMembers[memberID].BorrowedBooks {
		if n != l.allBooks[bookID] {
			updatedBorrowedBooks = append(updatedBorrowedBooks, n)
		}
	}

	member.BorrowedBooks = updatedBorrowedBooks
	l.allMembers[memberID] = member

	return nil
}

func (l *Library) ListAvailableBooks () []Book {
	available := []Book{}

	for _, n := range l.allBooks {
		if n.Status == "Available" {
			available = append(available, n)
		}
	}

	return available
}

func (l *Library) ListBorrowedBooks (memberID int) []Book {
	return l.allMembers[memberID].BorrowedBooks
}