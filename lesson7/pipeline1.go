package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		naturals <- 4
		close(naturals) // close the channel after sending data
	}()

	go func() {
		for x := range naturals { // range will exit the loop once the channel is closed
			squares <- x * x
		}
		close(squares) // close the squares channel when done
	}()

	for sq := range squares { // range will exit the loop once the squares channel is closed
		fmt.Println(sq)
	}
}

