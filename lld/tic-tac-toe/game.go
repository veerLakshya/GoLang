package tictactoe

import (
	"fmt"
	"sync"
)

type Game struct {
	board              *Board
	players            [2]*Player
	currentPlayerIndex int
	status             GameStatus
	winningStrategies  []WinningStrategy
	observers          []GameObserver
	mu                 sync.Mutex
}

func NewGame(player1, player2 *Player, boardSize int) (*Game, error) {
	if player1.Symbol() == player2.Symbol() {
		return nil, fmt.Errorf("Players must have different symbols")
	}
	g := &Game{
		players:            [2]*Player{player1, player2},
		board:              NewBoard(boardSize),
		currentPlayerIndex: 0,
		status:             StatusInProgress,
		observers:          make([]GameObserver, 0),
	}
	g.winningStrategies = g.initializeStrategies()
	return g, nil
}

func (g *Game) initializeStrategies() []WinningStrategy {
	return []WinningStrategy{
		&RowWinningStrategy{},
		&ColumnWinningStrategy{},
		&DiagonalWinningStrategy{},
	}
}

func (g *Game) MakeMove(row, col int) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.status != StatusInProgress {
		return NewInvalidMoveError("Game is already over!")
	}

	empty, err := g.board.IsCellEmpty(row, col)
	if err != nil {
		return err
	}
	if !empty {
		return NewInvalidMoveError("Cell (%d, %d) is already occupied", row, col)
	}
	currentPlayer := g.players[g.currentPlayerIndex]
	g.board.PlaceSymbol(row, col, currentPlayer.Symbol())

	if g.checkWin(row, col, currentPlayer.Symbol()) {
		if currentPlayer.Symbol() == SymbolX {
			g.status = StatusWinnerX
		} else {
			g.status = StatusWinnerO
		}
		g.notifyObservers()
		return nil
	}
	if g.board.IsFull() {
		g.status = StatusDraw
		g.notifyObservers()
		return nil
	}
	g.currentPlayerIndex = (g.currentPlayerIndex + 1) % 2
	return nil
}

func (g *Game) checkWin(row, col int, symbol Symbol) bool {
	for _, strategy := range g.winningStrategies {
		if strategy.CheckWin(g.board, row, col, symbol) {
			return true
		}
	}
	return false
}

func (g *Game) AddObserver(observer GameObserver) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.observers = append(g.observers, observer)
}
func (g *Game) notifyObservers() {
	for _, observer := range g.observers {
		observer.Update(g)
	}
}

func (g *Game) Board() *Board         { return g.board }
func (g *Game) CurrentPlaer() *Player { return g.players[g.currentPlayerIndex] }
func (g *Game) Status() GameStatus    { return g.status }

func (g *Game) Winner() *Player {
	switch g.status {
	case StatusWinnerX:
		if g.players[0].Symbol() == SymbolX {
			return g.players[0]
		}
		return g.players[1]
	case StatusWinnerO:
		if g.players[0].Symbol() == SymbolX {
			return g.players[0]
		}
		return g.players[1]
	default:
		return nil
	}
}

func (g *Game) PrintBoard() {
	g.board.PrintBoard()
}
