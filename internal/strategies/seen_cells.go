package strategies

import (
	"reflect"

	"github.com/alltilla/sudoku-solver/internal/sudoku"
)

func getAllCellsSeeing(grid *sudoku.Grid, cell *sudoku.Cell) [27]*sudoku.Cell {
	var cells [27]*sudoku.Cell

	if cells_in_row, err := grid.GetCellsInRow(cell.GetRowId()); err != nil {
		panic(err.Error())
	} else if cells_in_column, err := grid.GetCellsInColumn(cell.GetColumnId()); err != nil {
		panic(err.Error())
	} else if cells_in_box, err := grid.GetCellsInBox(cell.GetBoxId()); err != nil {
		panic(err.Error())
	} else {
		for i, cell := range cells_in_row {
			cells[i] = cell
		}
		for i, cell := range cells_in_column {
			cells[i+9] = cell
		}
		for i, cell := range cells_in_box {
			cells[i+18] = cell
		}
	}

	return cells
}

func SeenCells(grid *sudoku.Grid) (bool, error) {
	changed := false

	for _, cell := range grid.GetAllCells() {
		if cell.GetValue() != sudoku.Empty {
			continue
		}

		for _, other_cell := range getAllCellsSeeing(grid, cell) {
			if other_cell.GetValue() == sudoku.Empty {
				continue
			}

			old_pencil_marks := cell.GetPencilMarks()

			if err := cell.RemovePencilMark(other_cell.GetValue()); err != nil {
				panic(err.Error())
			}

			if !changed && !reflect.DeepEqual(cell.GetPencilMarks(), old_pencil_marks) {
				changed = true
			}
		}
	}

	return changed, nil
}
