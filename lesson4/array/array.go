package array

import "fmt"

func PrintArray() {
	var a [3]int = [...]int{1, 2, 3}

	fmt.Println(a, len(a), cap(a))

	var b [3]int = a
	fmt.Println(b)

}
