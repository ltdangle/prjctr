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
	id     int
	gameCh chan round
}
type guess struct {
	roundId  int
	playerId int
	number   int
}

func main() {
	var players []*player
	guesses := make(chan guess)
	roundTimer := make(chan round)

	// Init players.
	for i := 0; i < 5; i++ {
		p := &player{id: i, gameCh: make(chan round)}
		players = append(players, p)
		go playerGoroutine(p, guesses)
	}
	go roundGenerator(players, roundTimer)
	go roundReferee(players, guesses, roundTimer)
	time.Sleep(90 * time.Second)
}
func roundReferee(players []*player, guesses chan guess, roundTimer chan round) {
	for {
		select {
		case guess := <-guesses:
			fmt.Printf("\nroundReferee received: [round: %d, player: %d, number: %d]", guess.roundId, guess.playerId, guess.number)
		case round := <-roundTimer:
			fmt.Printf("\nroundReferee received new round notification for round %d", round.id)
		}
	}
}

func roundGenerator(players []*player, roundTimer chan round) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	roundCount := 0
	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("\nNew round at: %v", t)

			round := round{id: roundCount}

			for _, player := range players {
				player.gameCh <- round
			}

			roundTimer <- round

			roundCount++
		}
	}
}

func playerGoroutine(p *player, guesses chan guess) {
	fmt.Printf("\nPlayer %d is ready.", p.id)
	for {
		select {
		case round := <-p.gameCh:
			fmt.Printf("\nPlayer %d received new round number %d.", p.id, round.id)
			guess := guess{roundId: round.id, playerId: p.id, number: rand.Intn(5)}
			time.Sleep(time.Duration(guess.number) * time.Second)
			guesses <- guess
			fmt.Printf("\nPlayer %d for round number %d sent guess %d", p.id, round.id, guess.number)
		}
	}
}
