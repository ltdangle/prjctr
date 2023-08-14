package main

import (
	"fmt"
	library "hw19/library"
)

func main() {
	// Create lib.
	manager := library.NewManager("Steve")

	book1 := library.NewBook("Little Prince")
	book2 := library.NewBook("Some other book")
	book3 := library.NewBook("Yet another book")

	bookshelf1 := library.NewBookshelf("bookshelf1", 2)
	bookshelf2 := library.NewBookshelf("bookshelf2", 10)

	db := library.NewDb()

	lib := library.NewLibrary(db)
	lib.SetManager(manager)
	lib.AddBookshelf(bookshelf1)
	lib.AddBookshelf(bookshelf2)
	_ = lib.AddBook(book1)
	_ = lib.AddBook(book2)
	_ = lib.AddBook(book3)

	// Get our book from lib.
	myBook, err := lib.CheckoutBook("Little Prince")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(myBook)
}
