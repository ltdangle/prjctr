package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"
	"time"
)

type round struct {
	id     int
	number int
}
type player struct {
	id     int
	gameCh chan round
}
type guess struct {
	roundId   int
	playerId  int
	number    int
	guessedOn time.Time
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

	playedRounds := make(map[int][]guess)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nReferee is shutting down....")
			return
		case guess := <-guessChan:
			playedRounds[guess.roundId] = append(playedRounds[guess.roundId], guess)
			fmt.Printf("\nReferee received: [round: %d, player: %d, number: %d]", guess.roundId, guess.playerId, guess.number)
			fmt.Printf("\nPlayed rounds: %v", playedRounds)
		case round := <-refChan:
			fmt.Printf("\nReferee received new round notification for round %d", round.id)

			// Skip winner calculation on the first round notification
			if round.id == 0 {
				continue
			}

			// Calculate winner for prev. round.
			var winners []int
			prevRound := round.id - 1
			// Collect all winners
			for _, guess := range playedRounds[prevRound] {
				if round.number == guess.number {
					winners = append(winners, guess.playerId)
				}
			}

			w := winner{roundId: prevRound, playerId: -1}

			
			if len(winners) == 1 { // One winner
				w.playerId = winners[0]
			} else if len(winners) > 1 { // Several winners - find who send the winning guess first
				sort.Slice(playedRounds[prevRound], func(i, j int) bool {
					return playedRounds[prevRound][i].guessedOn.Before(playedRounds[prevRound][j].guessedOn)
				})
				w.playerId = winners[0]
			}

			fmt.Printf("\nThe winners for round %d is player %d !", prevRound, w.playerId)

			winnerCh <- winner{roundId: prevRound, playerId: w.playerId}

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

		round := round{id: roundCount, number: rand.Intn(4)}

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
			guess := guess{roundId: round.id, playerId: p.id, number: rand.Intn(4), guessedOn: time.Now()}
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
