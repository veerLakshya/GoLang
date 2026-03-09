package tictactoe

type Symbol int

const (
	SymbolX Symbol = iota
	SymbolO
	SymbolEmpty
)

func (s Symbol) DisplayChar() byte {
	switch s {
	case SymbolX:
		return 'X'
	case SymbolO:
		return 'O'
	default:
		return '_'
	}
}

func (s Symbol) String() string {
	switch s {
	case SymbolX:
		return "X"
	case SymbolO:
		return "O"
	default:
		return "_"
	}
}
