package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lib := Library{
		name:      "Abrhot",
		location:  "Addis Ababa",
		bookCount: map[int]int{},
		availableBooks: map[int]Book{
			// 1: {
			// 	ID:     1,
			// 	Title:  "The Go Programming Language",
			// 	Author: "Alan Donovan",
			// 	Status: "Available",
			// },
			// 2: {
			// 	ID:     2,
			// 	Title:  "Clean Code",
			// 	Author: "Robert C. Martin",
			// 	Status: "Borrowed",
			// },
			// 3: {
			// 	ID:     3,
			// 	Title:  "The Pragmatic Programmer",
			// 	Author: "Andrew Hunt",
			// 	Status: "Available",
			// },
			// 4: {
			// 	ID:     4,
			// 	Title:  "Introduction to Algorithms",
			// 	Author: "Thomas H. Cormen",
			// 	Status: "Borrowed",
			// },
			// 5: {
			// 	ID:     5,
			// 	Title:  "Head First Go",
			// 	Author: "Jay McGavren",
			// 	Status: "Available",
			// },
		},
		member: map[int]Member{},
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
			book := bookInput(scanner)
			lib.AddBook(book)
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
