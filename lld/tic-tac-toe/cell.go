package tictactoe

type Cell struct {
	symbol Symbol
}

func NewCell() *Cell {
	return &Cell{
		symbol: SymbolEmpty,
	}
}

func (c *Cell) Symbol() Symbol {
	return c.symbol
}

func (c *Cell) SetSymbol(s Symbol) {
	c.symbol = s
}

func (c *Cell) IsEmpty() bool {
	return c.symbol == SymbolEmpty
}
