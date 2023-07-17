package main

func main() {

  e:=NewEditor()
  e.AddFormatter(NewNoSpacesFormatter(NewUpperCaseWordsFormatter(NewText("some text here"))))
}
