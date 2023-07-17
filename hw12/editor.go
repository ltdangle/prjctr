package main

import "fmt"

type Editor struct {
	Text       Text
	Stats      map[string]TextStat
	Formatters map[string]Formatter
}

func NewEditor(text Text) *Editor {
	return &Editor{
		Text:       text,
		Stats:      make(map[string]TextStat),
		Formatters: make(map[string]Formatter),
	}
}

func (e *Editor) AddStatAction(s TextStat) {
	key := fmt.Sprintf("s%d", len(e.Stats)+1)
	e.Stats[key] = s
}

func (e *Editor) AddFormatter(f Formatter) {
	key := fmt.Sprintf("s%d", len(e.Formatters)+1)
	e.Formatters[key] = f
}

func (e *Editor) Menu() string {
	var formatters, stats string
	for key := range e.Formatters {
		formatters += fmt.Sprintf("%s: %s", key, "desc")
	}
	for key := range e.Stats {
		stats += fmt.Sprintf("%s: %s", key, "desc")
	}

	return fmt.Sprintf(`
Text:
%s

Text statistics:
%s

Text formatters:
%s
`, e.Text.Format(), stats, formatters)
}
