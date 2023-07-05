package main

import (
	"fmt"
	"time"
	// "time"
)
type round struct {
	id int
}
type player struct {
	playerNumber int
	gameCh       chan round
}

func main() {
	var players []*player
	// Init players.
	for i := 0; i < 5; i++ {
		p := &player{playerNumber: i, gameCh: make(chan round)}
		players = append(players, p)
		go playerGoroutine(p)
	}
	go gameLoop(players)
	time.Sleep(9 * time.Second)
}

func gameLoop(players []*player) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	roundCount:=0
	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("\nNew round at: %v", t)
			for _, player := range players {
				player.gameCh <-round{id:roundCount} 
			}
			roundCount++
		}
	}
}

func playerGoroutine(p *player) {
	fmt.Printf("\nPlayer %d is ready.", p.playerNumber)
	for {
		select {
		case round:=<-p.gameCh:
			fmt.Printf("\nPlayer %d received new round number %d.", p.playerNumber, round.id)
		}
	}
}
