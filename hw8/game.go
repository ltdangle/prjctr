package main

import (
	"fmt"
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
	// time.Sleep(2*time.Second)
}

func playerGoroutine(p *player) {
	fmt.Printf("\nPlayer %d is ready.", p.playerNumber)
}
