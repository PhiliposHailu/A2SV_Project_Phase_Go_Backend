package main

import (
	"bufio"
	"fmt"
	"strconv"
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
	ListAvailableBooks()
	ListBorrowedBooks(memberID int)
}

type Library struct {
	name           string
	location       string
	bookCount      map[int]int
	availableBooks map[int]Book
	member         map[int]Member
}

// DISPLAYS INFORMAION
func displayAvailableBooks(listOfBooks map[int]Book, count map[int]int) {
	fmt.Printf("| %-8v | %-8v |\n", "Book", "Count")
	for key := range listOfBooks {
		fmt.Printf("| %-8v | %-8v |\n", listOfBooks[key].Title, count[key])
	}
}

func diplayBorrowedBooks(name string, booksBorrowedList []Book) {
	fmt.Printf("List of borrowed books by: %v\n", name)
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

func bookInput(scanner *bufio.Scanner, id int) Book {

	fmt.Print("Enter Book Title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Enter Book Author: ")
	scanner.Scan()
	author := scanner.Text()

	bookStatus := Available

	return Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: bookStatus,
	}

}

func memberIdInput(scanner *bufio.Scanner) int {
	fmt.Print("Enter your member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())

	return memberID
}

func fullMemberInfoInput(scanner *bufio.Scanner, memberId int) Member {
	fmt.Printf("Enter your name: ")
	scanner.Scan()
	name := scanner.Text()

	return Member{
		ID:            memberId,
		Name:          name,
		BorrowedBooks: []Book{},
	}

}

// LIBRARY INTERFACE IMPLEMENATION
func (l Library) AddBook(b Book) {
	_, exists := l.availableBooks[b.ID]
	if !exists {
		l.availableBooks[b.ID] = b
	}
	l.bookCount[b.ID]++
}

func (l Library) RemoveBook(bookID int) {
	if l.bookCount[bookID] >= 1 {
		book := l.availableBooks[bookID]
		l.bookCount[bookID]--
		// if l.bookCount[bookID] == 0 {
		// 	l.availableBooks[bookID] = 
		// }
		fmt.Printf("Book %v removed succesfully", book.Title)
		return
	}

	fmt.Printf("Book with id %v Not Found! \n", bookID)
}

func (l Library) BorrowBook(scanner *bufio.Scanner, bookID int, memberID int) {
	if l.bookCount[bookID] >= 1 {
		// register a member
		_, exists := l.member[memberID]

		var member Member
		if !exists {
			member = fullMemberInfoInput(scanner, memberID)
			l.member[memberID] = member

		}
		member = l.member[memberID]

		// update the book list of the member
		book := l.availableBooks[bookID]
		member.BorrowedBooks = append(member.BorrowedBooks, book) // make a copy
		l.member[memberID] = member                               // resign the copy back to our dicionary
		l.bookCount[bookID]--

		fmt.Printf("Book '%s' borrowed successfuly", book.Title)
		return
	}

	fmt.Printf("Book with ID: '%v' \nNot Available/Not Found ... :3\n", bookID)

}

func (l Library) ReturnBook(bookID int, memberID int) {
	found := false
	idx := -1
	for i, val := range l.member[memberID].BorrowedBooks {
		if bookID == val.ID {
			found = true
			idx = i
			break
		}
	}

	if found {
		l.bookCount[bookID]++
		mem := l.member[memberID]
		mem.BorrowedBooks = append(mem.BorrowedBooks[:idx], mem.BorrowedBooks[idx+1:]...)
		l.member[memberID] = mem
		fmt.Printf("Book '%v' successfuly returned.", l.availableBooks[bookID].Title)
	} else {
		fmt.Printf("Book Not Found")
	}
}

func (l Library) ListAvailableBooks() {
	displayAvailableBooks(l.availableBooks, l.bookCount)
}

func (l Library) ListBorrowedBooks(memberId int) {
	name := l.member[memberId].Name
	books := l.member[memberId].BorrowedBooks
	diplayBorrowedBooks(name, books)
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
