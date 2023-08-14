package library

import "errors"

type BookRecord struct {
	bookShelfId BookshelfId
	BookTitle   BookTitle
}
type Db struct {
	storage []*BookRecord
}

func NewDb() *Db {
	return &Db{}
}

func (db *Db) IndexBook(bookshelfId BookshelfId, bookTitle BookTitle) {
	db.storage = append(db.storage, &BookRecord{bookShelfId: bookshelfId, BookTitle: bookTitle})
}

func (db *Db) Find(title BookTitle) (BookshelfId, error) {
	for _, bookRecord := range db.storage {
		if bookRecord.BookTitle == title {
			return bookRecord.bookShelfId, nil
		}
	}
	return "", errors.New("could not find book location in db")
}
