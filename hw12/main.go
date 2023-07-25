package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const statKey = "stat"
const formatKey = "format"

// Receives text via stdin. Formats text and shows text statistics according cli flags.
// Example usage: `echo "your text here" | go run . -stat=s2 -format=lower,no-spaces,upper`
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

	// Parse cli flags.
	statAction := flag.String(statKey, "", StatHelp(e))
	formatAction := flag.String(formatKey, "", FormatterHelp(e))
	flag.Parse()

	// Formatters.
	formatterKeys := strings.Split(*formatAction, ",")
	var format Formatter = e.Format(formatterKeys)
	fmt.Printf("\n%s:\n%s", format.Description(), format.Format())

	// Text stats.
	if *statAction != "" {
		stat := e.Stats[*statAction]
		fmt.Printf("\n%s: %d\n", stat.Description(), stat.Count())
	}
}

// Cli help entry for formatters.
func FormatterHelp(e *Editor) string {
	var formatters string
	for key, _ := range e.Formatters {
		formatters += fmt.Sprintf("-%s=%s\n", formatKey, key)
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
