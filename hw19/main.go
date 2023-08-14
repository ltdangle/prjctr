package main

import (
	"fmt"
	"github.com/google/uuid"
	db2 "hw19/db"
	library "hw19/library"
)

func main() {

	customer := library.NewCustomer("You")
	manager := library.NewManager("Steve")

	uuid1 := _randomUuid()
	book1 := library.NewBook(uuid1, "Little Prince")
	book2 := library.NewBook(_randomUuid(), "Some other book")
	book3 := library.NewBook(_randomUuid(), "Yet another book")

	bookshelf1 := library.NewBookshelf("bookshelf1", 2)
	bookshelf2 := library.NewBookshelf("bookshelf2", 10)

	db := db2.NewDb()

	lib := library.NewLibrary(db)
	lib.SetManager(manager)
	lib.AddBookshelf(bookshelf1)
	lib.AddBookshelf(bookshelf2)
	_ = lib.AddBook(book1)
	_ = lib.AddBook(book2)
	_ = lib.AddBook(book3)

	// Get our book from lib.
	myBook, err := lib.CheckoutBook("Little Prince", customer)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(myBook)
}

func _randomUuid() library.BookUuid {
	return library.BookUuid(uuid.New().String())
}
