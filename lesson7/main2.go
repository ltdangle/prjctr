package main

import "fmt"

func foo(c chan int, someValue int) {
	c <- someValue
}
func main() {
	fooVal := make(chan int)

	go foo(fooVal, 1)
	go foo(fooVal, 2)
	go foo(fooVal, 3)

	v1 := <-fooVal
	v2 := <-fooVal
	v3 := <-fooVal

	fmt.Println(v1, v2, v3)

}
