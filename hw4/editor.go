package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Editor.
type editor struct {
	text          []string
	matchingLines []int
}

func (e *editor) readIntput() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter some text. When you're done, press Ctrl+D (or Ctrl+Z in Windows):")

	for scanner.Scan() {
		e.text = append(e.text, scanner.Text())
	}
}
func (e *editor) readSearchString() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nSearch string: ")
	scanner.Scan()
	search := scanner.Text()
	for i, line := range e.text {
		if strings.Contains(line, search) {
			e.matchingLines = append(e.matchingLines, i)

		}
	}
}
func (e *editor) printSearchResults() {
	if len(e.matchingLines) == 0 {
		fmt.Println("No matches found!")
		return
	}
	fmt.Print("\nMatches found:\n")
	for _, lineNumber := range e.matchingLines {
		fmt.Printf("%d: %s\n", lineNumber, e.text[lineNumber])
	}
}
