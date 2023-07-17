package main

import "fmt"

func main() {
  t:=NewText("some text here")
  e:=NewEditor(t)
  e.AddStatAction(NewWhiteSpaceCounter())
  e.AddFormatter(NewNoSpacesFormatter(NewUpperCaseWordsFormatter(t)))
  fmt.Println(e.Menu())
}
