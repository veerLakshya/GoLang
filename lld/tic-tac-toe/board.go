package tictactoe

import (
	"fmt"
	"strings"
)

type Board struct {
	grid [][]*Cell
	size int
}

func NewBoard(size int) *Board {
	b := &Board{
		size: size,
		grid: make([][]*Cell, size),
	}
	b.initializeBoard()
	return b
}

func (b *Board) initializeBoard() {
	for i := 0; i < b.size; i++ {
		b.grid[i] = make([]*Cell, b.size)
		for j := 0; j < b.size; j++ {
			b.grid[i][j] = NewCell()
		}
	}
}

func (b *Board) PlaceSymbol(row, col int, symbol Symbol) error {
	if err := b.validatePosition(row, col); err != nil {
		return err
	}
	b.grid[row][col].SetSymbol(symbol)
	return nil
}

func (b *Board) IsCellEmpty(row, col int) (bool, error) {
	err := b.validatePosition(row, col)
	if err != nil {
		return false, err
	}
	return b.grid[row][col].IsEmpty(), err
}

func (b *Board) IsFull() bool {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.grid[i][j].IsEmpty() {
				return false
			}
		}
	}

	return true
}

func (b *Board) GetCell(row, col int) (*Cell, error) {
	err := b.validatePosition(row, col)
	if err != nil {
		return nil, err
	}
	return b.grid[row][col], nil
}

func (b *Board) Size() int {
	return b.size
}

func (b *Board) validatePosition(row, col int) error {
	if row < 0 || col < 0 || row >= b.size || col >= b.size {
		return NewInvalidMoveError("Position (%d %d) is out of bonds", row, col)
	}
	return nil
}

func (b *Board) PrintBoard() {
	fmt.Println()
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			fmt.Printf(" %c ", b.grid[i][j].Symbol().DisplayChar())
			if j < b.size-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < b.size-1 {
			fmt.Println(strings.Repeat("_", b.size*4-1))
		}
	}
	fmt.Println()
}
