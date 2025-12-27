package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lib := Library{
		name:           "Abrhot",
		location:       "Addis Ababa",
		bookCount:      map[int]int{},
		availableBooks: map[int]Book{},
		member:         map[int]Member{},
	}

	fmt.Println("Welcome to the Library Mangement System :)")
	exitProgram := false
	for !exitProgram {
		diplay()
		pickedNum := choice()
		switch pickedNum {
		case 0:
			exitProgram = true
		case 1:
			bookId := bookIdInput()
			_, exists := lib.availableBooks[bookId]
			
			if !exists {
				book := bookInput(scanner, bookId)
				lib.AddBook(book)
			} else {
				lib.bookCount[bookId]++
			}
			
		case 2:
			id := bookIdInput()
			lib.RemoveBook(id)
		case 3:
			bookId := bookIdInput()
			memId := memberIdInput(scanner)
			lib.BorrowBook(scanner, bookId, memId)
		case 4:
			bookId := bookIdInput()
			memId := memberIdInput(scanner)
			lib.ReturnBook(bookId, memId)
		case 5:
			lib.ListAvailableBooks()
		case 6:
			memId := memberIdInput(scanner)
			lib.ListBorrowedBooks(memId)
		}
	}

}
