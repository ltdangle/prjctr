package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
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

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	// Register shutdown function.
	shutDownListener(cancel, &wg)

	// Players send guesses this channel.
	guessChan := make(chan guess)
	// Referee channel.
	refChan := make(chan round)
	// Winner channel.
	winnerCh := make(chan winner)

	// Init players and player gorutines.
	for i := 0; i < 5; i++ {
		p := &player{id: i, gameCh: make(chan round)}
		players = append(players, p)
		wg.Add(1)
		go playerRtn(ctx, p, guessChan, &wg)
	}
	
	// Init round generator.
	wg.Add(1)
	go roundGeneratorRtn(ctx, players, refChan, &wg)

	// Init referee gorutine.
	wg.Add(1)
	go refereeRtn(ctx, players, guessChan, refChan, winnerCh, &wg)

	// Poll for winner.
	for {
		select {
		case winner := <-winnerCh:
			fmt.Printf("\n------------------\nThe winner for round %d is player %d !\n------------------", winner.roundId, winner.playerId)
		default:
		}
	}
}

func refereeRtn(ctx context.Context, players []*player, guessChan chan guess, refChan chan round, winnerCh chan winner, wg *sync.WaitGroup) {
	defer wg.Done()

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

func roundGeneratorRtn(ctx context.Context, players []*player, refChan chan round, wg *sync.WaitGroup) {
	defer wg.Done()

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
			fmt.Printf("\nRound generator is shutting down....")
			return
		case t := <-ticker.C:
			roundFn(t)
		}
	}
}

func playerRtn(ctx context.Context, p *player, guesses chan guess, wg *sync.WaitGroup) {
	defer wg.Done()

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

func shutDownListener(cancel func(), wg *sync.WaitGroup) {

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
		wg.Wait()
		os.Exit(0)
	}()
}
