package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
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
	ctx, cancel := context.WithCancel(context.Background())

	shutDownListener(cancel)

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
		go playerGoroutine(ctx, p, guessChan)
	}
	go roundGenerator(ctx, players, refChan)
	go roundReferee(ctx, players, guessChan, refChan, winnerCh)

	for {
		select {
		case winner := <-winnerCh:
			fmt.Printf("\n------------------\nThe winner for round %d is player %d !\n------------------", winner.roundId, winner.playerId)
		default:
		}
	}
}

func roundReferee(ctx context.Context, players []*player, guessChan chan guess, refChan chan round, winnerCh chan winner) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nReferee is shutting down....")
			return
		case guess := <-guessChan:
			fmt.Printf("\nReferee received: [round: %d, player: %d, number: %d]", guess.roundId, guess.playerId, guess.number)
		case round := <-refChan:
			fmt.Printf("\nReferee received new round notification for round %d", round.id)
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

func roundGenerator(ctx context.Context, players []*player, refChan chan round) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	roundCount := 0

	roundFn := func(t time.Time) {
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

	// Run periodically from now on.
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Round generator is shutting down....")
			return
		case t := <-ticker.C:
			roundFn(t)
		}
	}
}

func playerGoroutine(ctx context.Context, p *player, guesses chan guess) {
	fmt.Printf("\nPlayer %d is ready.", p.id)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nGorutine for player %d is shutting down...", p.id)
			return
		case round := <-p.gameCh:
			guess := guess{roundId: round.id, playerId: p.id, number: rand.Intn(5)}
			time.Sleep(time.Duration(guess.number) * time.Second)
			guesses <- guess
			fmt.Printf("\nPlayer %d for round number %d sent guess %d", p.id, round.id, guess.number)
		}
	}
}

func shutDownListener(cancel func()) {
	// Create a channel to receive OS signals
	sigs := make(chan os.Signal, 1)

	// Register the channel to receive SIGINT signals (CTRL-C)
	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		fmt.Println("Exiting......")
		cancel()
		// TODO
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()
}
