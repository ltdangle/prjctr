package main

import (
	"fmt"
	"math/rand"
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
	guesses:=make(chan int)
	// Init players.
	for i := 0; i < 5; i++ {
		p := &player{playerNumber: i, gameCh: make(chan round)}
		players = append(players, p)
		go playerGoroutine(p, guesses)
	}
	go roundGenerator(players)
	go roundReferee(players,guesses)
	time.Sleep(90 * time.Second)
}
func roundReferee(players []*player, guesses chan int) {
	for {
		select {
		case guess := <-guesses:
			fmt.Printf("\nroundReferee received: %d", guess)
		}
	}
}

func roundGenerator(players []*player) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	roundCount := 0
	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("\nNew round at: %v", t)
			for _, player := range players {
				player.gameCh <- round{id: roundCount}
			}
			roundCount++
		}
	}
}

func playerGoroutine(p *player, guesses chan int) {
	fmt.Printf("\nPlayer %d is ready.", p.playerNumber)
	for {
		select {
		case round := <-p.gameCh:
			fmt.Printf("\nPlayer %d received new round number %d.", p.playerNumber, round.id)
			guess:=rand.Intn(5)
			time.Sleep(time.Duration(guess)*time.Second)
			guesses<-guess
			fmt.Printf("\nPlayer %d for round number %d sent guess %d", p.playerNumber, round.id, guess)
		}
	}
}
