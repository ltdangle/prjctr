package main

import "fmt"

func main() {
  t:=NewText("some text here")
  e:=NewEditor(t)
  e.AddStatAction(NewWhiteSpaceCounter())
  e.AddFormatter(NewNoSpacesFormatter(NewUpperCaseWordsFormatter(t)))
  e.AddFormatter(NewUpperCaseWordsFormatter(NewNoSpacesFormatter(t)))
  fmt.Println(e.Menu())
}
