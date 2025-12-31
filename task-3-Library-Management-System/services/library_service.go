package services

import (
	"fmt"

	"github.com/philipos/library/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
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
	availableBooks map[int]models.Book
	member         map[int]models.Member
}

func (l *Library) CheckAvailability(bookId int) bool {
	_, exists := l.availableBooks[bookId]
	return exists
}

func (l *Library) UpdateBookCount(bookId int, sign string) {
	if sign == "+" {
		l.bookCount[bookId]++
	} else {
		l.bookCount[bookId]--
	}
}

func NewLibrary(newName string, newLocation string) *Library {
	lib := Library{
		name:           newName,
		location:       newLocation,
		bookCount:      map[int]int{},
		availableBooks: map[int]models.Book{},
		member:         map[int]models.Member{},
	}

	return &lib
}

// DISPLAYS INFORMAION
func displayAvailableBooks(listOfBooks map[int]models.Book, count map[int]int) {
	fmt.Println()
	fmt.Printf("| %-16v | %-16v | %-16v |\n", "#", "Book", "Count")
	row := 1
	for key := range listOfBooks {
		if count[key] > 0 {
			fmt.Printf("| %-16v | %-16v | %-16v |\n", row, listOfBooks[key].Title, count[key])
			row++
		}
	}
}

func diplayBorrowedBooks(name string, booksBorrowedList []models.Book) {
	fmt.Println()
	fmt.Printf("List of borrowed books by: %v\n", name)
	fmt.Printf("| %-16v | %-16v | %-16v |\n", "#", "Title", "Author")
	row := 1
	for i := range len(booksBorrowedList) {
		book := booksBorrowedList[i]
		fmt.Printf("| %-16v | %-16v | %-16v |\n", row, book.Title, book.Author)
	}
}

// LIBRARY INTERFACE IMPLEMENATION
func (l *Library) AddBook(b models.Book) {
	_, exists := l.availableBooks[b.ID]
	if !exists {
		l.availableBooks[b.ID] = b
	}
	l.bookCount[b.ID]++
}

func (l *Library) RemoveBook(bookID int) {
	if l.bookCount[bookID] >= 1 {
		book := l.availableBooks[bookID]
		l.bookCount[bookID]--

		fmt.Printf("Book %v removed succesfully", book.Title)
		return
	}

	fmt.Printf("Book with id %v Not Found! \n", bookID)
}

func (l *Library) MemberExists(memberId int) bool {
	_, exists := l.member[memberId]
	return exists
}

func (l *Library) ResgisterNewMember(member models.Member, memberId int) {
	l.member[memberId] = member
}

func (l *Library) BorrowBook(bookID int, memberID int) error {

	if l.bookCount[bookID] >= 1 {

		_, exists := l.member[memberID]
		if exists {
			if len(l.member[memberID].BorrowedBooks) > 2 {
				fmt.Println("Book Limit Reached. \nPlease return books to take more!")
				return nil
			}
		}

		member := l.member[memberID]

		// update the book list of the member
		book := l.availableBooks[bookID]
		member.BorrowedBooks = append(member.BorrowedBooks, book) // make a copy
		l.member[memberID] = member                               // resign the copy back to our dicionary
		l.bookCount[bookID]--

		fmt.Printf("Book '%s' borrowed successfuly", book.Title)
		return nil
	}

	fmt.Printf("Book with ID: '%v' \nNot Available/Not Found ... :3\n", bookID)
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) {
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

func (l *Library) ListAvailableBooks() {
	displayAvailableBooks(l.availableBooks, l.bookCount)
}

func (l *Library) ListBorrowedBooks(memberId int) {
	name := l.member[memberId].Name
	books := l.member[memberId].BorrowedBooks
	diplayBorrowedBooks(name, books)
}
