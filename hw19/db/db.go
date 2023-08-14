package db

import (
	"errors"
	"github.com/google/uuid"
	"hw19/library"
)

type RecordUuid string

type Db struct {
	storage map[RecordUuid]*library.Book
}

func NewDb() *Db {
	return &Db{
		storage: make(map[RecordUuid]*library.Book),
	}
}

func (db *Db) IndexBook(book *library.Book) {
	RecordUuid := RecordUuid(uuid.New().String())
	db.storage[RecordUuid] = book
}

func (db *Db) Find(title library.BookTitle) (*library.Book, error) {
	for _, book := range db.storage {
		if book.Title == title {
			return book, nil
		}
	}
	return nil, errors.New("could not find book location in db")
}
