package editor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Editor.
type Editor struct {
	text          []string
	matchingLines []int
}

func (e *Editor) ReadIntput() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter some text. When you're done, press Ctrl+D (or Ctrl+Z in Windows):")

	for scanner.Scan() {
		e.text = append(e.text, scanner.Text())
	}
}
func (e *Editor) ReadSearchString() {
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
func (e *Editor) PrintSearchResults() {
	if len(e.matchingLines) == 0 {
		fmt.Println("No matches found!")
		return
	}
	fmt.Print("\nMatches found:\n")
	for _, lineNumber := range e.matchingLines {
		fmt.Printf("%d: %s\n", lineNumber, e.text[lineNumber])
	}
}
func (e *Editor) Run() {
	e.ReadIntput()
	e.ReadSearchString()
	e.PrintSearchResults()
}
