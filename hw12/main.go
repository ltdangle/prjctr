package main

import "fmt"

func main() {

  e:=NewEditor()
  e.AddFormatter(NewNoSpacesFormatter(NewUpperCaseWordsFormatter(NewText("some text here"))))
  fmt.Println(e.Menu())
}
