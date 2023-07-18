package main

import "regexp"

// TextStat statistics interface.
type TextStat interface {
	// Count Statistic action.
	Count() int
	// Description Statistic description.
	Description() string
}

// SpaceCounterLoop whitespace counter implementation using loop.
type SpaceCounterLoop struct{ text Text }

func NewWhiteSpaceCounter(text Text) *SpaceCounterLoop {
	return &SpaceCounterLoop{text: text}
}

func (c *SpaceCounterLoop) Count() int {
	counter := 0
	for i := 0; i < len(c.text.Text); i++ {
		if c.text.Text[i:i+1] == " " {
			counter++
		}
	}
	return counter
}

func (c *SpaceCounterLoop) Description() string {
	return "Counts spaces via loop."
}

// SpaceCounterRegex whitespace counter implementation using regex.
type SpaceCounterRegex struct{ text Text }

func NewWhiteSpaceCounterRegex(text Text) *SpaceCounterRegex {
	return &SpaceCounterRegex{text: text}
}

func (c *SpaceCounterRegex) Count() int {
	re := regexp.MustCompile(` `)
	matches := re.FindAllString(c.text.Text, -1)

	return len(matches)
}
func (c *SpaceCounterRegex) Description() string {
	return "Counts spaces via regex."
}
