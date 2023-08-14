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

func (db *Db) findRecordUuid(book library.Book) (RecordUuid, error) {
	for recordUuid, b := range db.storage {
		if b.Id == book.Id {
			return recordUuid, nil
		}
	}
	return "", nil
}

func (db *Db) Update(book library.Book) error {
	b, err := db.Find(book.Title)
	if err != nil {
		return err
	}

	recordUuid, err := db.findRecordUuid(book)
	if err != nil {
		return err
	}

	// Update key value.
	db.storage[recordUuid] = b

	return nil
}
