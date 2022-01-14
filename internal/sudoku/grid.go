package sudoku

import (
	"fmt"
	"math"
	"strconv"
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

func parseValueFromChar(char byte) (int, error) {
	if char == ' ' {
		return Empty, nil
	} else {
		return strconv.Atoi(string(char))
	}
}

func (g *Grid) LoadValuesFromString(str string) error {
	for i := 0; i < 9*(9+1); i++ {
		row := int(math.Floor(float64(i)/10) + 1)
		column := (i % 10) + 1
		char := str[i]

		if column == 10 {
			if char != '\n' {
				return fmt.Errorf("missing NL at position %d", i)
			}
			continue
		}

		cell, err := g.GetCell(row, column)
		if err != nil {
			return err
		}

		value, err := parseValueFromChar(char)
		if err != nil {
			return err
		}

		err = cell.SetValue(value)
		if err != nil {
			return err
		}
	}

	return nil
}
