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
	member         map[int]Member
}

// DISPLAYS INFORMAION
func displayAvailableBooks(listOfBooks map[int]Book) {
	fmt.Printf("|%v |%-8v|", "Book", "Count")
	for key, count := range listOfBooks {
		fmt.Printf("|%v | %-8v|", key, count)
	}
}

func diplayBorrowedBooks(name string, booksBorrowedList []Book) {
	fmt.Printf("List of borrowed books by %v: \n", name)
	fmt.Printf("| %-8v | %-8v | %-8v |\n", "#", "Title", "Author")
	for i := range len(booksBorrowedList) {
		book := booksBorrowedList[i]
		fmt.Printf("| %-8v | %-8v | %-8v |\n", book.ID, book.Title, book.Author)
	}
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

func fullMemberInfoInput(memberId int) Member {
	var name string

	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)

	return Member{
		ID:            memberId,
		Name:          name,
		BorrowedBooks: []Book{},
	}

}

// LIBRARY INTERFACE IMPLEMENATION
func (l Library) AddBook(b Book) {
	l.availableBooks[b.ID] = b
	l.bookCount[b.ID]++
}

func (l Library) RemoveBook(bookID int) {
	if l.bookCount[bookID] >= 1 {
		book := l.availableBooks[bookID]
		l.bookCount[bookID]--
		fmt.Printf("Book %v removed succesfully", book.Title)
		return
	}

	fmt.Printf("Book with id %v Not Found! \n", bookID)
}

func (l Library) BorrowBook(bookID int, memberID int) {
	if l.bookCount[bookID] >= 1 {
		// register a member
		memberID := memberIdInput()
		_, exists := l.member[memberID]

		var member Member
		if !exists {
			member = fullMemberInfoInput(memberID)
		} else {
			member = l.member[memberID]
		}

		// update the book list of the member
		book := l.availableBooks[bookID]
		member.BorrowedBooks = append(member.BorrowedBooks, book)
		l.bookCount[bookID]--

		fmt.Printf("Book '%s' borrowed successfuly", book.Title)
		return
	}

	fmt.Printf("Book with ID: '%v' \nNot Available/Not Found ... :3\n", bookID)

}

func (l Library) ReturnBook(bookID int, memberID int) {
	l.bookCount[bookID]++
	fmt.Printf("Book '%v' successfuly returned.", l.availableBooks[bookID].Title)

}

// Display Options
func diplay() {
	fmt.Printf(`Choose One:
	1, to Add a Book
	2, to Remove a Book
	3, to Boorow a Book
	4, to Return a Book
	5, to List all Available Books
	6, to list all Borrowed Books
	0, to Exit
	`)
}

func choice() int {
	var num int
	fmt.Scan(&num)
	for num < 0 || num > 6 {
		fmt.Print("Enter a Valid number \n(Note that the number you pick must be with in the range of 0 - 6)")
		fmt.Scan(&num)
	}
	return num
}
