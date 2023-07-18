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
	key := fmt.Sprintf("f%d", len(e.Formatters)+1)
	e.Formatters[key] = f
}
