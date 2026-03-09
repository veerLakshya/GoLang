package tictactoe

type RowWinningStrategy struct{}

func (s *RowWinningStrategy) CheckWin(b *Board, row, col int, symbol Symbol) bool {
	size := b.Size()

	for c := 0; c < size; c++ {
		cell, _ := b.GetCell(row, c)
		if cell.Symbol() != symbol {
			return false
		}
	}
	return true
}

type ColumnWinningStrategy struct{}

func (s *ColumnWinningStrategy) CheckWin(b *Board, row, col int, symbol Symbol) bool {
	size := b.Size()

	for r := 0; r < size; r++ {
		cell, _ := b.GetCell(r, col)
		if cell.Symbol() != symbol {
			return false
		}
	}
	return true
}

type DiagonalWinningStrategy struct{}

func (s *DiagonalWinningStrategy) CheckWin(b *Board, row, col int, symbol Symbol) bool {
	size := b.Size()

	mainDiagonalWin := true
	for i := 0; i < size; i++ {
		cell, _ := b.GetCell(i, i)
		if cell.Symbol() != symbol {
			mainDiagonalWin = false
			break
		}
	}
	if mainDiagonalWin {
		return true
	}

	for i := 0; i < size; i++ {
		cell, _ := b.GetCell(i, size-i-1)
		if cell.Symbol() != symbol {
			return false
		}
	}
	return true
}
