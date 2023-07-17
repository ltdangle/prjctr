package main

import "regexp"

// Text statistics interface.
type TextStat interface {
	// Statistic action.
	Count(text string) int
}

// Whitespace counter implementation using loop.
type SpaceCounterLoop struct{}

func NewWhiteSpaceCounter() *SpaceCounterLoop {
	return &SpaceCounterLoop{}
}

func (c *SpaceCounterLoop) Count(text string) int {
	counter := 0
	for i := 0; i < len(text); i++ {
		if text[i:i+1] == " " {
			counter++
		}
	}
	return counter
}

// Whitespace counter implementation using regex.
type SpaceCounterRegex struct{}

func NewWhiteSpaceCounterRegex() *SpaceCounterRegex {
	return &SpaceCounterRegex{}
}

func (c *SpaceCounterRegex) Count(text string) int {
	re := regexp.MustCompile(` `)
	matches := re.FindAllString(text, -1)

	return len(matches)

}
