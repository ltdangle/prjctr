package main

import "fmt"

type Editor struct {
	Text       Text
	Stats      map[string]TextStat
	Formatters map[string]func(t Formatter) Formatter
}

func NewEditor(text Text) *Editor {
	e := Editor{
		Text:       text,
		Stats:      make(map[string]TextStat),
		Formatters: make(map[string]func(t Formatter) Formatter),
	}

	// Init available formatters.
	e.Formatters = map[string]func(t Formatter) Formatter{
		"no-spaces": func(t Formatter) Formatter {
			return NewNoSpacesFormatter(t)
		},
		"upper": func(t Formatter) Formatter {
			return NewUpperCaseWordsFormatter(t)
		},
		"lower": func(t Formatter) Formatter {
			return NewLowerCaseWordsFormatter(t)
		},
	}
	return &e
}

// Returns format decorator chain.
func (e *Editor) Format(formatterKeys []string) Formatter {
	var f Formatter = e.Text
	for _, formatter := range formatterKeys {
		f = e.Formatters[formatter](f)
	}
	return f
}

func (e *Editor) AddStatAction(s TextStat) {
	key := fmt.Sprintf("s%d", len(e.Stats)+1)
	e.Stats[key] = s
}
