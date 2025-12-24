package main

import "fmt"

func main () {

	// BORROW a book 
	memberID := memberIdInput()
	bookID := bookIdInput()
	fmt.Printf("%v, %v \n", memberID, bookID)
	
}
