package main

import (
	"fmt"
)

// Створити дві горутини і обʼєднати їх каналами (string). Деяка строка мусить пройти дві горутини,
// де кожна її допише, і ми отримаємо її на виході
func main() {
	var name string
	c := make(chan string)
	go addSuffix(c, " Bob")
	go addSuffix(c, " Mark")
	c <- "John"
	name = <-c
	fmt.Println(name)

}
func addSuffix(c chan string, suffix string) {
	str := <-c
	c <- str + suffix
}
