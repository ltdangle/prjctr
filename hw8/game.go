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
	go gameLoop(players)
	time.Sleep(9 * time.Second)
}

func gameLoop(players []*player) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("\nNew round at: %v", t)
			for _, player := range players {
				player.gameCh <- ""
			}
		}
	}
}

func playerGoroutine(p *player) {
	fmt.Printf("\nPlayer %d is ready.", p.playerNumber)
	for {
		select {
		case <-p.gameCh:
			fmt.Printf("\nPlayer %d received new round.", p.playerNumber)
		}
	}
}
