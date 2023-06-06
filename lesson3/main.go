package main

import (
	"fmt"
	"math/rand"
)

// "fmt"
// "ifelse/ifelse"
// "math/rand"
// "time"

func main() {
	// shouldPrint := true
	// if comparisonResult := ifelse.IsABigger(3, 4); comparisonResult && shouldPrint {
	// 	fmt.Println("branch one")
	// } else {
	// 	fmt.Println("branch two")
	// }

	// rand.Seed(time.Now().UnixNano())
	// switch i := rand.Intn(10); i % 2 {
	// case 0, 5:
	// 	fmt.Println(i, 0)
	// case 1:
	// 	fmt.Println(i, 1)
	// default:
	// 	fmt.Println("Default case")
	// }

	// for n := 0; n != 2; n = rand.Intn(10) {
	// 	fmt.Println(n)
	// }
	// fmt.Println("...End")

	var n int
	for n != 2 {
		n = rand.Intn(10)
		fmt.Println(n)
	}
}
