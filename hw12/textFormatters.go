package main

import "strings"

// Formatter interface.
type Formatter interface {
	// TODO: rename to Text()
	Format() string
	Description() string
}

// Text struct (implements Formatter interface).
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
	return strings.ReplaceAll(f.Formatter.Format(), " ", "")
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
	words := strings.Split(f.Formatter.Format(), " ")
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return strings.Join(words, " ")
}
func (f *UpperCaseWordsFormatter) Description() string {
	return f.Formatter.Description() + "Words to uppercase. "
}

// LowerCaseWordsFormatter.
type LowerCaseWordsFormatter struct {
	Formatter Formatter
}

func NewLowerCaseWordsFormatter(formatter Formatter) *LowerCaseWordsFormatter {
	return &LowerCaseWordsFormatter{formatter}
}

func (f *LowerCaseWordsFormatter) Format() string {
	words := strings.Split(f.Formatter.Format(), " ")
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, " ")
}
func (f *LowerCaseWordsFormatter) Description() string {
	return f.Formatter.Description() + "Words to lowercase. "
}
