package tictactoe

import (
	"fmt"
	"sync"
)

type ScoreBoard struct {
	scores map[string]int
	mu     sync.RWMutex
}

func NewScoreboard() *ScoreBoard {
	return &ScoreBoard{
		scores: make(map[string]int),
	}
}

func (sb *ScoreBoard) Update(game *Game) {
	winner := game.Winner()
	if winner != nil {
		sb.RecordWin(winner)
		fmt.Printf("Scoreboard updated: %s wins!\n", winner.Name())
	}
}

func (sb *ScoreBoard) RecordWin(player *Player) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	sb.scores[player.Name()]++
}

func (sb *ScoreBoard) GetScore(playerName string) int {
	sb.mu.RLock()
	defer sb.mu.RUnlock()
	return sb.scores[playerName]
}

func (sb *ScoreBoard) PrintScoreboard() {
	sb.mu.RLock()
	defer sb.mu.RUnlock()

	fmt.Println("\n=========Scoreboard=========")
	if len(sb.scores) == 0 {
		fmt.Println("No games played yet.")
	} else {
		for name, score := range sb.scores {
			fmt.Printf("%s: %d wins\n", name, score)
		}
	}
	fmt.Println("\n===========================")
}
