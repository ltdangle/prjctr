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

	// Вcі слова з великої літери, що закінчуються на голосну
	regex := regexp.MustCompile(`\p{Lu}\p{Ll}*[аеєиіїоуюя]`)

	matches := regex.FindAllString(s, -1)

	// Print the matches.
	for _, match := range matches {
		fmt.Println(match)
	}
}
