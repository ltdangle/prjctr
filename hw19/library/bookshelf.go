package library

import (
	"errors"
)

type BookshelfId string
type BookTitle string
type Bookshelf struct {
	id       BookshelfId
	capacity int
	books    []*Book
}

func NewBookshelf(id BookshelfId, capacity int) *Bookshelf {
	return &Bookshelf{
		id:       id,
		capacity: capacity,
	}
}

// addBook adds book to a bookshelf.
func (s *Bookshelf) addBook(book *Book) error {
	// Check that there is space in the bookshelf.
	if len(s.books) == s.capacity {
		return errors.New("no more space on the bookshelf")
	}

	s.books = append(s.books, book)

	return nil
}

// Book returns book from the bookshelf.
func (s *Bookshelf) Book(title BookTitle) (*Book, bool) {
	for location, book := range s.books {
		if book.Title == title {
			return s.books[location], true
		}
	}
	return nil, false
}
