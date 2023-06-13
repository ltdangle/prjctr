package main

import "hw4/editor"

func main() {
	e := &editor.Editor{}
	e.ReadIntput()
	e.ReadSearchString()
	e.PrintSearchResults()
}
