package main

import (
	"fmt"
	"time"
	// "time"
)

type player struct {
	playerNumber int
	gameCh       chan interface{}
}

func main() {
	var players []*player
	// Init players.
	for i := 0; i < 5; i++ {
		p := &player{playerNumber: i, gameCh: make(chan interface{})}
		players = append(players, p)
		go playerGoroutine(p)
	}
	go gameLoop()
	time.Sleep(9*time.Second)
}

func gameLoop() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("\nCurrent time: %v", t)
		}
	}
}
func playerGoroutine(p *player) {
	fmt.Printf("\nPlayer %d is ready.", p.playerNumber)
}
