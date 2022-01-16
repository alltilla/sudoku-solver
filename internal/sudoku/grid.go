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

func (g *Grid) GetCell(row int, column int) (*Cell, error) {
	err := CheckRowAndColumnValidity(row, column)
	if err != nil {
		return nil, err
	}
	return g.cells[row-1][column-1], nil
}

func (g *Grid) GetCellsInRow(row int) ([9]*Cell, error) {
	cells := [9]*Cell{nil, nil, nil, nil, nil, nil, nil, nil, nil}

	for column := 1; column <= 9; column++ {
		cell, err := g.GetCell(row, column)
		if err != nil {
			return cells, err
		}

		cells[column-1] = cell
	}

	return cells, nil
}

func (g *Grid) GetCellsInColumn(column int) ([9]*Cell, error) {
	cells := [9]*Cell{nil, nil, nil, nil, nil, nil, nil, nil, nil}

	for row := 1; row <= 9; row++ {
		cell, err := g.GetCell(row, column)
		if err != nil {
			return cells, err
		}

		cells[row-1] = cell
	}

	return cells, nil
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
