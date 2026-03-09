package tictactoe

type WinningStrategy interface {
	CheckWin(board *Board, row, col int, symbol Symbol) bool
}

type GameObserver interface {
	Update(game *Game)
}
