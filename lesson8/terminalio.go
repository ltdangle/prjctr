package main

import (
	"fmt"
	"log"
	"time"

	"github.com/eiannone/keyboard"
)

var input string

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	// Channel to signal when to stop the clock goroutine
	stopClock := make(chan bool)

	// Start a goroutine that prints the time every second
	go func() {
		for {
			select {
			case <-stopClock:
				return
			default:
				fmt.Print("\033[H\033[2J")
				fmt.Printf("\rCurrent time: %s", time.Now().Format(time.RFC1123))
				fmt.Printf("\nInput: %s", input)
				time.Sleep(1 * time.Second)
			}
		}
	}()

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

	// Main loop for reading keys
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		// If ESC is pressed, send a signal to stop the clock goroutine and exit
		if key == keyboard.KeyEsc {
			stopClock <- true
			break
		}
		input += string(char)
		fmt.Printf("\rYou pressed: %c ( %X)\n", char, key)
	}
}

