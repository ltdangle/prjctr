package main

import "fmt"

type Book struct {
	title         string
	author        string
	numberOfPages int
}
type BookShelf struct {
	firstBook Book
}

func (b *Book) GetPages() int {
	return b.numberOfPages
}

func (b *Book) AddPage() {
	b.numberOfPages++
}

func main() {
	var b Book
	b.AddPage()
	fmt.Println(b.GetPages())
}
