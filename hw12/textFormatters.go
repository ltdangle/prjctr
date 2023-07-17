package main

import "strings"

type Formatter interface {
	// TODO: rename to Text()
	Format() string
	Description() string
}

// Text struct.
type Text struct {
	Text string
}

func NewText(text string) Text {
	return Text{Text: text}
}

func (t Text) Format() string {
	return t.Text
}

func (t Text) Description() string {
	return "Original text. "
}

// NoSpacesFormatter.
type NoSpacesFormatter struct {
	Formatter Formatter
}

func NewNoSpacesFormatter(formatter Formatter) *NoSpacesFormatter {
	return &NoSpacesFormatter{formatter}
}

func (f *NoSpacesFormatter) Format() string {
	return strings.ReplaceAll(f.Format(), " ", "")
}
func (f *NoSpacesFormatter) Description() string {
	return f.Formatter.Description() + "No spaces. "
}

// UpperCaseWordsFormatter.
type UpperCaseWordsFormatter struct {
	Formatter Formatter
}

func NewUpperCaseWordsFormatter(formatter Formatter) *UpperCaseWordsFormatter {
	return &UpperCaseWordsFormatter{formatter}
}

func (f *UpperCaseWordsFormatter) Format() string {
	words := strings.Split(f.Format(), " ")
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return strings.Join(words, " ")
}
func (f *UpperCaseWordsFormatter) Description() string {
	return f.Formatter.Description() + "Words to uppercase. "
}
