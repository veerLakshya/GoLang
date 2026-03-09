package tictactoe

import (
	"fmt"
	"sync"
)

type TicTacToeSystem struct {
	scoreboard  *ScoreBoard
	currentGame *Game
}

var (
	instance *TicTacToeSystem
	once     sync.Once
	resetMu  sync.Mutex
)

func GetInstance() *TicTacToeSystem {
	once.Do(func() {
		instance = &TicTacToeSystem{
			scoreboard: NewScoreboard(),
		}
	})
	return instance
}

func (s *TicTacToeSystem) CreateGame(player1, player2 *Player) (*Game, error) {
	game, err := NewGame(player1, player2, 3)
	if err != nil {
		return nil, err
	}
	s.currentGame = game
	s.currentGame.AddObserver(s.scoreboard)
	fmt.Printf("New game started: %s vs %s\n", player1.Name(), player2.Name())
	return s.currentGame, nil
}

func (s *TicTacToeSystem) MakeMove(player *Player, row, col int) error {
	if s.currentGame == nil {
		return fmt.Errorf("no active game: call CreateGame first")
	}
	fmt.Printf("%s plays at (%d, %d)\n", player.Name(), row, col)
	if err := s.currentGame.MakeMove(row, col); err != nil {
		return err
	}
	s.currentGame.PrintBoard()
	return nil
}

func (s *TicTacToeSystem) GameStatus() (GameStatus, error) {
	if s.currentGame == nil {
		return StatusInProgress, fmt.Errorf("no active game")
	}
	return s.currentGame.Status(), nil
}

func ResetInstance() {
	resetMu.Lock()
	defer resetMu.Unlock()
	instance = nil
	once = sync.Once{}
}
