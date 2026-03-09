package tictactoe

type GameStatus int

const (
	StatusInProgress GameStatus = iota
	StatusWinnerX
	StatusWinnerO
	StatusDraw
)

func (gs GameStatus) String() string {
	switch gs {
	case StatusInProgress:
		return "IN_PROGRESS"
	case StatusWinnerX:
		return "WINNER_X"
	case StatusWinnerO:
		return "WINNER_O"
	case StatusDraw:
		return "DRAW"
	default:
		return "UNKNOWN"
	}
}
