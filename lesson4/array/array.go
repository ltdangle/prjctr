package array

import "fmt"

func printArray(a [3]int) {
	fmt.Println(a, len(a), cap(a))
}
func PrintArray() {
	var a [3]int = [...]int{1, 2, 3}
	printArray(a)

	var b [3]int = a
	b[0] = 5
	printArray(a)
	printArray(b)
}
