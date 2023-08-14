package library

// Book struct.
type Book struct {
	Title        BookTitle
	BookShelfId  BookshelfId
	CheckedOutBy UserId
}

func NewBook(title BookTitle) *Book {
	return &Book{Title: title}
}
