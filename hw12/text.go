package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {

	bytes, err := os.ReadFile("1689007676028_text.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := string(bytes)

	regex := regexp.MustCompile(`\b[бвгґджзклмнпрстфхцчшщ]\S*\s`)
	matches := regex.FindAllString(s, -1)

	// Print the matches.
	for _, match := range matches {
		fmt.Println(match)
	}
}
