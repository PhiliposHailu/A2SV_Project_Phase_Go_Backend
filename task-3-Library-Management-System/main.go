package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/philipos/library/controller"
	"github.com/philipos/library/services"

)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lib := services.NewLibrary("Abrhot", "Addis Ababa")

	fmt.Println("Welcome to the Library Mangement System :)")
	exitProgram := false
	for !exitProgram {
		controller.Diplay()
		pickedNum := controller.Choice()
		switch pickedNum {
		case 0:
			exitProgram = true
		case 1:
			bookId := controller.BookIdInput()
			available := lib.CheckAvailability(bookId)

			if !available {
				book := controller.BookInput(scanner, bookId)
				lib.AddBook(book)
			} else {
				lib.UpdateBookCount(bookId, "+")
			}

		case 2:
			id := controller.BookIdInput()
			lib.RemoveBook(id)
		case 3:
			bookId := controller.BookIdInput()
			memId := controller.MemberIdInput(scanner)

			// Register new user 
			if !lib.MemberExists(memId) {
				member := controller.FullMemberInfoInput(scanner, memId)
				lib.ResgisterNewMember(member, memId)
			}
			lib.BorrowBook(bookId, memId)
		case 4:
			bookId := controller.BookIdInput()
			memId := controller.MemberIdInput(scanner)
			lib.ReturnBook(bookId, memId)
		case 5:
			lib.ListAvailableBooks()
		case 6:
			memId := controller.MemberIdInput(scanner)
			lib.ListBorrowedBooks(memId)
		}
	}

}
