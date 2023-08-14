package library

// Book struct.
type Book struct {
	title BookTitle
}

func NewBook(title BookTitle) *Book {
	return &Book{title: title}
}
