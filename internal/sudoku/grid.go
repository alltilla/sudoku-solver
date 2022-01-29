package sudoku

import (
	"strings"
)

type Grid struct {
	cells        [9][9]*Cell
	cells_by_box [9][]*Cell
}

type Set struct {
	Orientation string
	Index       int
	Cells       [9]*Cell
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

func (g *Grid) GetSets() [27]*Set {
	var sets [27]*Set

	for i := 0; i < 9; i++ {
		row := i + 1
		if cells_in_row, err := g.GetCellsInRow(row); err != nil {
			panic(err.Error())
		} else {
			sets[i] = &Set{"row", row, cells_in_row}
		}
	}

	for i := 0; i < 9; i++ {
		column := i + 1
		if cells_in_column, err := g.GetCellsInColumn(column); err != nil {
			panic(err.Error())
		} else {
			sets[i+9] = &Set{"column", column, cells_in_column}
		}
	}

	for i := 0; i < 9; i++ {
		box := i + 1
		if cells_in_box, err := g.GetCellsInBox(box); err != nil {
			panic(err.Error())
		} else {
			sets[i+18] = &Set{"box", box, cells_in_box}
		}
	}

	return sets
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

func (g *Grid) Equals(other *Grid) bool {
	for row := 1; row <= 9; row++ {
		for column := 1; column <= 9; column++ {
			cell, err := g.GetCell(row, column)
			if err != nil {
				panic(err.Error())
			}

			other_cell, err := other.GetCell(row, column)
			if err != nil {
				panic(err.Error())
			}

			if !cell.Equals(other_cell) {
				return false
			}
		}
	}

	return true
}

func (g *Grid) makeFullyEmpty() {
	for _, cell := range g.GetAllCells() {
		if err := cell.SetValue(Empty); err != nil {
			panic(err.Error())
		}

		if err := cell.RemovePencilMarks([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}); err != nil {
			panic(err.Error())
		}
	}
}

func (g *Grid) LoadPrettyString(pretty_string string) error {
	pretty_string_trimmed := strings.TrimSpace(pretty_string)

	if err := validateFrame(pretty_string_trimmed); err != nil {
		return err
	}

	g.makeFullyEmpty()

	if err := g.deserializeGridPrettyString(pretty_string_trimmed); err != nil {
		return err
	}

	return nil
}
