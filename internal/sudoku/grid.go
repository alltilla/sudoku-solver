package sudoku

import (
	"fmt"
	"math"
	"strings"
)

type Grid struct {
	cells        [9][9]*Cell
	cells_by_box [9][]*Cell
}

func NewGrid() *Grid {
	g := Grid{}

	for row := 1; row <= 9; row++ {
		for column := 1; column <= 9; column++ {
			cell, err := NewCell(row, column)

			if err != nil {
				panic("failed to create cell")
			}

			g.cells[row-1][column-1] = cell

			box_id := cell.GetBoxId()
			g.cells_by_box[box_id-1] = append(g.cells_by_box[box_id-1], cell)
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

func (g *Grid) GetCellsInBox(box int) ([9]*Cell, error) {
	cells := [9]*Cell{nil, nil, nil, nil, nil, nil, nil, nil, nil}

	err := CheckBoxValidity(box)
	if err != nil {
		return cells, err
	}

	copy(cells[:], g.cells_by_box[box-1])

	return cells, nil
}

func (g *Grid) GetAllCells() [81]*Cell {
	var cells [81]*Cell

	i := 0
	for row := 1; row <= 9; row++ {
		for column := 1; column <= 9; column++ {
			cell, err := g.GetCell(row, column)
			if err != nil {
				panic("failed to get cell")
			}

			cells[i] = cell
			i++
		}
	}

	return cells
}

func (g *Grid) LoadValuesFromString(str string) error {
	for i := 0; i < 9*(9+1); i++ {
		row := int(math.Floor(float64(i)/10) + 1)
		column := (i % 10) + 1
		cell_value_str := string(str[i])

		if column == 10 {
			if cell_value_str != "\n" {
				return fmt.Errorf("missing NL at position %d", i)
			}
			continue
		}

		cell, err := g.GetCell(row, column)
		if err != nil {
			return err
		}

		err = cell.LoadValueFromString(cell_value_str)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Grid) LoadPencilMarksFromString(str string) error {
	pencil_marks := strings.FieldsFunc(str, func(r rune) bool { return r == '|' || r == '\n' })

	if len(pencil_marks) != 81 {
		return fmt.Errorf("unexpected amount of pencil marks: %d", len(pencil_marks))
	}

	for i := 0; i < 9*9; i++ {
		row := int(math.Floor(float64(i)/9) + 1)
		column := (i % 9) + 1

		cell, err := g.GetCell(row, column)
		if err != nil {
			panic("failed to get cell")
		}

		err = cell.LoadPencilMarksFromString(pencil_marks[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Grid) Equals(other *Grid) bool {
	for row := 1; row <= 9; row++ {
		for column := 1; column <= 9; column++ {
			cell, err := g.GetCell(row, column)
			if err != nil {
				panic("failed to get cell")
			}

			other_cell, err := other.GetCell(row, column)
			if err != nil {
				panic("failed to get cell")
			}

			if !cell.Equals(other_cell) {
				return false
			}
		}
	}

	return true
}
