package sudoku

import (
	"fmt"
)

type Grid struct {
	cells [9][9]*Cell
}

func NewGrid() *Grid {
	g := Grid{}

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			g.cells[row][column] = NewCell()
		}
	}

	return &g
}

func checkRowAndColumnValidity(row int, column int) error {
	if row > 9 {
		return fmt.Errorf("row is larger than 9: %d", row)
	}
	if row < 1 {
		return fmt.Errorf("row is smaller than 1: %d", row)
	}
	if column > 9 {
		return fmt.Errorf("column is larger than 9: %d", column)
	}
	if column < 1 {
		return fmt.Errorf("column is smaller than 1: %d", column)
	}

	return nil
}

func (g *Grid) GetCell(row int, column int) (*Cell, error) {
	err := checkRowAndColumnValidity(row, column)
	if err != nil {
		return nil, err
	}
	return g.cells[row-1][column-1], nil
}
