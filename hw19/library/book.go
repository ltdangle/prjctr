package library

type BookUuid string

// Book struct.
type Book struct {
	Id               BookUuid
	Title            BookTitle
	BookShelfId      BookshelfId
	IsCheckedOut     bool
	LastCheckedOutBy UserId
}

func NewBook(bookUuid BookUuid, title BookTitle) *Book {
	return &Book{
		Id:           bookUuid,
		Title:        title,
		IsCheckedOut: false,
	}
}
