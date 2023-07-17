package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	bytes, err := os.ReadFile("1689007675141_numbers.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := string(bytes)

	// The regular expression pattern
	pattern := `\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}`

	// Compile the pattern
	re := regexp.MustCompile(pattern)

	// Find all matches in the string
	matches := re.FindAllString(s, -1)

	// Print the matches
	for _, match := range matches {
		fmt.Println(match)
	}
}
