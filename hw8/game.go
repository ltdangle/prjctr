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
type winner struct {
	roundId  int
	playerId int
}

func main() {
	var players []*player
	guessChan := make(chan guess)
	// Referee channel.
	refChan := make(chan round)
	// Winner channel.
	winnerCh := make(chan winner)

	// Init players.
	for i := 0; i < 5; i++ {
		p := &player{id: i, gameCh: make(chan round)}
		players = append(players, p)
		go playerGoroutine(p, guessChan)
	}
	go roundGenerator(players, refChan)
	go roundReferee(players, guessChan, refChan, winnerCh)

	for {
		select {
		case winner := <-winnerCh:
			fmt.Printf("\n------------------\nThe winner for round %d is player %d !\n------------------", winner.roundId, winner.playerId)
		default:
		}
	}
	// TODO: handle shutdown
}

func roundReferee(players []*player, guessChan chan guess, refChan chan round, winnerCh chan winner) {
	for {
		select {
		case guess := <-guessChan:
			fmt.Printf("\nroundReferee received: [round: %d, player: %d, number: %d]", guess.roundId, guess.playerId, guess.number)
		case round := <-refChan:
			fmt.Printf("\nroundReferee received new round notification for round %d", round.id)
			// Randomly calculate winner for prev. round.
			if round.id > 0 {
				prevRound := round.id - 1
				winnerId := rand.Intn(len(players))
				fmt.Printf("\nThe winner for %d is player %d !", prevRound, winnerId)
				winnerCh <- winner{roundId: prevRound, playerId: winnerId}
			}
		}
	}
}

func roundGenerator(players []*player, refChan chan round) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	roundCount := 0

	roundFn:=func(t time.Time){
			fmt.Printf("\nNew round at: %v", t)

			round := round{id: roundCount}

			for _, player := range players {
				player.gameCh <- round
			}

			refChan <- round

			roundCount++
	}

	// Run on start, don't wait for ticker on first run.
	roundFn(time.Now())
	for {
		select {
		case t := <-ticker.C:
		roundFn(t)
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
