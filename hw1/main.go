package main

import "fmt"

func main() {
	fmt.Println("Steve Jobs biography:")
	fmt.Println("---------------------")
	fmt.Printf("Steve Jobs was born on %v in San Francisco, California.\n", "February 24, 1955")
	fmt.Printf("He co-founded %s with Steve Wozniak in %d.\n", "Apple Inc.", 1976)
	fmt.Printf("He was also involved with other well-known companies such as %s and %s.\n", "Pixar", "NeXT")
	fmt.Printf("Jobs passed away on %v due to a rare form of pancreatic cancer.\n", "October 5, 2011")
	fmt.Printf("His influence on the technology industry and creative design continues to be felt today.\n")
	fmt.Printf("At the time of his death, Steve Jobs' net worth was approximately $%f billion.\n", 8.3)
}
