package tictactoe

import "fmt"

type Player struct {
	name   string
	symbol Symbol
}

func NewPlayer(name string, symbol Symbol) (*Player, error) {
	if name == "" {
		return nil, fmt.Errorf("player name cannot be empty")
	}
	if symbol == SymbolEmpty {
		return nil, fmt.Errorf("player symbol cannot be empty")
	}
	return &Player{
		name:   name,
		symbol: symbol,
	}, nil
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Symbol() Symbol {
	return p.symbol
}

func (p *Player) String() string {
	return fmt.Sprintf("%s (%c)", p.name, p.symbol.DisplayChar())
}
