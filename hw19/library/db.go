package library

import (
	"errors"
	"github.com/google/uuid"
)

type RecordUuid string

type BookRecord struct {
	bookShelfId BookshelfId
	BookTitle   BookTitle
}
type Db struct {
	storage map[RecordUuid]*BookRecord
}

func NewDb() *Db {
	return &Db{
		storage: make(map[RecordUuid]*BookRecord),
	}
}

func (db *Db) IndexBook(bookshelfId BookshelfId, bookTitle BookTitle) {
	RecordUuid := RecordUuid(uuid.New().String())
	db.storage[RecordUuid] = &BookRecord{bookShelfId: bookshelfId, BookTitle: bookTitle}
}

func (db *Db) Find(title BookTitle) (BookshelfId, error) {
	for _, bookRecord := range db.storage {
		if bookRecord.BookTitle == title {
			return bookRecord.bookShelfId, nil
		}
	}
	return "", errors.New("could not find book location in db")
}
