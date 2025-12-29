package main

import (
	"bufio"
	"fmt"
	"os"
	// "github.com/philipos/library/models"
	// "github.com/philipos/library/services"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// finish up next itme
	// lib := services.Library{
	// 	name:           "Abrhot",
	// 	location:       "Addis Ababa",
	// 	bookCount:      map[int]int{},
	// 	availableBooks: map[int]models.Book{},
	// 	member:         map[int]models.Member{},
	// }

	fmt.Println("Welcome to the Library Mangement System :)")
	exitProgram := false
	// for !exitProgram {
	// 	controller.Diplay()
	// 	pickedNum := controller.Choice()
	// 	switch pickedNum {
	// 	case 0:
	// 		exitProgram = true
	// 	case 1:
	// 		bookId := controller.BookIdInput()
	// 		_, exists := lib.availableBooks[bookId]

	// 		if !exists {
	// 			book := controller.BookInput(scanner, bookId)
	// 			lib.AddBook(book)
	// 		} else {
	// 			lib.bookCount[bookId]++
	// 		}

	// 	case 2:
	// 		id := controller.BookIdInput()
	// 		lib.RemoveBook(id)
	// 	case 3:
	// 		bookId := controller.BookIdInput()
	// 		memId := controller.MemberIdInput(scanner)
	// 		lib.BorrowBook(scanner, bookId, memId)
	// 	case 4:
	// 		bookId := controller.BookIdInput()
	// 		memId := controller.MemberIdInput(scanner)
	// 		lib.ReturnBook(bookId, memId)
	// 	case 5:
	// 		lib.ListAvailableBooks()
	// 	case 6:
	// 		memId := controller.MemberIdInput(scanner)
	// 		lib.ListBorrowedBooks(memId)
	// 	}
	// }

}
