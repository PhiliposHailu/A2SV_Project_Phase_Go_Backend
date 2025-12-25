package main

import "fmt"

func main() {
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
		pickedNum := choice()
		switch pickedNum {
		case 0:
			exitProgram = true
		case 1:
			book := bookInput()
			lib.AddBook(book)
		case 2:
			id := bookIdInput()
			lib.RemoveBook(id)
		case 3:
			bookId := bookIdInput()
			memId := memberIdInput()
			lib.BorrowBook(bookId, memId)
		case 4:
			bookId := bookIdInput()
			memId := memberIdInput()
			lib.ReturnBook(bookId, memId)
		case 5:
			// lib.availableBooks()
		}
	}

}
