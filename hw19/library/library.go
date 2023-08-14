package library

import "errors"

// Library struct.
type Library struct {
	manager    *Manager
	bookshelfs map[BookshelfId]*Bookshelf
	db         IDb
}

// IDb database interface.
type IDb interface {
	IndexBook(book *Book)
	Find(title BookTitle) (*Book, error)
	Update(book *Book) error
}

// NewLibrary constructor.
func NewLibrary(db IDb) *Library {
	return &Library{
		bookshelfs: make(map[BookshelfId]*Bookshelf),
		db:         db,
	}
}

func (l *Library) SetManager(manager *Manager) {
	l.manager = manager
}

func (l *Library) AddBookshelf(bookshelf *Bookshelf) {
	l.bookshelfs[bookshelf.id] = bookshelf
}

func (l *Library) AddBook(book *Book) error {
	var bookshelfId BookshelfId
	// adds book to next available bookcase
	for _, bookshelf := range l.bookshelfs {
		err := bookshelf.addBook(book)
		if err != nil {
			continue
		}
		bookshelfId = bookshelf.id
		break
	}

	if bookshelfId == "" {
		return errors.New("could not add book to bookshelf")
	}

	// adds location and Title to the database
	book.BookShelfId = bookshelfId
	l.db.IndexBook(book)

	return nil
}

// CheckoutBook retrieves book from the library.
func (l *Library) CheckoutBook(title BookTitle, customer *Customer) (*Book, error) {
	b, err := l.db.Find(title)
	if err != nil {
		return nil, err
	}

	// Locate bookshelf.
	_, bookshelfExists := l.bookshelfs[b.BookShelfId]
	if !bookshelfExists {
		return nil, errors.New("could not find bookshelf")
	}

	// Locate book.
	book, bookExists := l.bookshelfs[b.BookShelfId].Book(title)
	if !bookExists {
		return nil, errors.New("could not find the book on the bookshelf")
	}

	// Update book status.
	book.IsCheckedOut = true
	book.LastCheckedOutBy = customer.id

	return book, nil
}

func (l *Library) ReturnBook(book *Book) error {
	foundBook, err := l.db.Find(book.Title)
	if err != nil {
		return err
	}

	if foundBook.Id != book.Id {
		return errors.New("trying to return wrong book")
	}

	foundBook.IsCheckedOut = false

	err = l.db.Update(book)
	if err != nil {
		return err
	}

	return nil
}
