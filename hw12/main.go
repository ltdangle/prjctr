package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

const statKey = "stat"
const formatKey = "format"

// Receives text via stdin. Formats text and shows text statistics according cli flags.
// Example usage: `echo "your text here" | go run . -stat=s2 -format=f2`
// Help: `echo "your text here" | go run . -help`
func main() {
	var text string

	// Read text from stdin.
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}
		text += input
	}

	// Assemble editor.
	t := NewText(text)
	e := NewEditor(t)
	e.AddStatAction(NewWhiteSpaceCounter(t))
	e.AddStatAction(NewWhiteSpaceCounterRegex(t))
	e.AddFormatter(NewNoSpacesFormatter(NewUpperCaseWordsFormatter(t)))
	e.AddFormatter(NewUpperCaseWordsFormatter(NewNoSpacesFormatter(t)))
	e.AddFormatter(NewLowerCaseWordsFormatter(t))

	// Parse cli flags.
	statAction := flag.String(statKey, "", StatHelp(e))
	formatAction := flag.String(formatKey, "", FormatterHelp(e))
	flag.Parse()

	// Run action according to cli flags.
	if *formatAction != "" {
		format := e.Formatters[*formatAction]
		fmt.Printf("\n%s:\n%s", format.Description(), format.Format())
	}
	if *statAction != "" {
		stat := e.Stats[*statAction]
		fmt.Printf("\n%s: %d\n", stat.Description(), stat.Count())
	}
}

// Cli help entry for formatters.
func FormatterHelp(e *Editor) string {
	var formatters string
	for key, formatter := range e.Formatters {
		formatters += fmt.Sprintf("-%s=%s -  %s\n", formatKey, key, formatter.Description())
	}
	return formatters
}

// Cli help entry for text stat options.
func StatHelp(e *Editor) string {
	var stats string
	for key, stat := range e.Stats {
		stats += fmt.Sprintf("-%s=%s - %s\n", statKey, key, stat.Description())
	}
	return stats
}
