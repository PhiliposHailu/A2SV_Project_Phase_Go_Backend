package controller

import (
	"bufio"
	"fmt"
	"strconv"
	"github.com/philipos/library/models"
)

// Display Options
func Diplay() {
	fmt.Println()
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

func Choice() int {
	var num int
	fmt.Scan(&num)
	for num < 0 || num > 6 {
		fmt.Print("Enter a Valid number \n(Note that the number you pick must be with in the range of 0 - 6)")
		fmt.Scan(&num)
	}
	return num
}


// USER INPUTS Functions
func BookIdInput() int {
	fmt.Print("Enter a book ID: ")
	var bookID int
	fmt.Scan(&bookID)

	return bookID
}

func BookInput(scanner *bufio.Scanner, id int) models.Book {

	fmt.Print("Enter Book Title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Enter Book Author: ")
	scanner.Scan()
	author := scanner.Text()

	var bookStatus models.BookStatus = models.Available

	return models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: bookStatus,
	}

}

func MemberIdInput(scanner *bufio.Scanner) int {
	fmt.Print("Enter your member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())

	return memberID
}

func FullMemberInfoInput(scanner *bufio.Scanner, memberId int) models.Member {
	fmt.Printf("Enter your name: ")
	scanner.Scan()
	name := scanner.Text()

	return models.Member{
		ID:            memberId,
		Name:          name,
		BorrowedBooks: []models.Book{},
	}

}