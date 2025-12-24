package main

import (
	"fmt"
)

type BookStatus string

const (
	Available BookStatus = "Available"
	Borrowed  BookStatus = "Borrowed"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Status BookStatus
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
	name           string
	location       string
	bookCount      map[int]int
	availableBooks map[int]Book
	borrowBooks    map[int]Book
	member         map[int]Member
}

// DISPLAYS INFORMAION
func displayAvailableBooks(listOfBooks map[int]Book) {
	fmt.Printf("|%v |%-8v|", "Book", "Count")
	for key, count := range listOfBooks {
		fmt.Printf("|%v | %-8v|", key, count)
	}
}

func diplayBorrowedBooks(memberID, listofBorrowers map[int]Book) {
	
}

// USER INPUTS Functions
func bookIdInput() int {
	fmt.Print("Enter a book ID: ")
	var bookID int
	fmt.Scan(&bookID)

	return bookID
}

func bookInput() Book {
	fmt.Print("Enter Book ID: ")
	var id int
	fmt.Scan(&id)

	fmt.Print("Enter Book Title: ")
	var title string
	fmt.Scan(&title)

	fmt.Print("Enter Book Author: ")
	var author string
	fmt.Scan(&author)

	bookStatus := Available

	return Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: bookStatus,
	}

}

func memberIdInput() int {
	fmt.Print("Enter your member ID: ")
	var memberID int
	fmt.Scan(&memberID)

	return memberID
}


// LIBRARY INTERFACE IMPLEMENATION
func (l Library) AddBook(b Book) {
	l.availableBooks[b.ID] = b
	l.bookCount[b.ID]++
}

func (l Library) RemoveBook(bookID int) {
	if l.bookCount[bookID] >= 1 {
		l.bookCount[bookID]--
		fmt.Printf("Book %v added succesfully", l.availableBooks[bookID].Title)
		return
	}

	fmt.Printf("Sorry Book with id %v Not Available in the first place! in order to remove\n", bookID)
}

func (l Library) BorrowBook(bookID int, memberID int) {
	if l.bookCount[bookID] >= 1 {
		// register a member
		_, exists := l.member[memberID]
		if !exists {
			l.member[memberID] = Member{
				ID:            1,
				Name:          "lala",
				BorrowedBooks: []Book{},
			}
		}

		book := l.availableBooks[bookID]
		m := l.member[memberID]
		m.BorrowedBooks = append(m.BorrowedBooks, book)

		l.bookCount[bookID]--

		fmt.Printf("Book %s borrowed successfuly", l.availableBooks[bookID].Title)
		return
	}

	fmt.Printf("book with id %v not available ... :3\n", bookID)
}

func (l Library) ReturnBook(bookID int, memberID int) {
	l.bookCount[bookID]++
	fmt.Printf("Book %v successfuly returned.", l.availableBooks[bookID])

}
