package strategies

import (
	"fmt"

	"github.com/alltilla/sudoku-solver/internal/sudoku"
)

func pencilMarksContainDigit(pencil_marks []int, digit int) bool {
	for _, pencil_mark := range pencil_marks {
		if pencil_mark == digit {
			return true
		}
	}

	return false
}

func findHiddenSingleInArrayOfCells(cells [9]*sudoku.Cell) (bool, error) {

	for digit := 1; digit <= 9; digit++ {
		var hidden_single_cell_candidate *sudoku.Cell = nil
		value_found := false
		multiple_candidates := false

		for _, cell := range cells {
			if cell.GetValue() == digit {
				value_found = true
				break
			}

			if pencilMarksContainDigit(cell.GetPencilMarks(), digit) {
				if hidden_single_cell_candidate != nil {
					multiple_candidates = true
					break
				}

				hidden_single_cell_candidate = cell
			}
		}

		if value_found || multiple_candidates {
			continue
		}

		if hidden_single_cell_candidate == nil {
			return false, fmt.Errorf("no possible cell to place digit: %d", digit)
		}

		if err := hidden_single_cell_candidate.SetValue(digit); err != nil {
			panic(err.Error())
		}
		return true, nil
	}

	return false, nil
}

func findHiddenSingleInRows(grid *sudoku.Grid) (bool, error) {
	for row := 1; row <= 9; row++ {
		cells_in_row, err := grid.GetCellsInRow(row)
		if err != nil {
			panic(err.Error())
		}

		if changed, err := findHiddenSingleInArrayOfCells(cells_in_row); err != nil {
			return changed, fmt.Errorf("error in row %d: %w", row, err)
		} else if changed {
			return changed, err
		}
	}

	return false, nil
}

func findHiddenSingleInColumns(grid *sudoku.Grid) (bool, error) {
	for column := 1; column <= 9; column++ {
		cells_in_column, err := grid.GetCellsInColumn(column)
		if err != nil {
			panic(err.Error())
		}

		if changed, err := findHiddenSingleInArrayOfCells(cells_in_column); err != nil {
			return changed, fmt.Errorf("error in column %d: %w", column, err)
		} else if changed {
			return changed, err
		}
	}

	return false, nil
}

func findHiddenSingleInBoxes(grid *sudoku.Grid) (bool, error) {
	for box := 1; box <= 9; box++ {
		cells_in_box, err := grid.GetCellsInBox(box)
		if err != nil {
			panic(err.Error())
		}

		if changed, err := findHiddenSingleInArrayOfCells(cells_in_box); err != nil {
			return changed, fmt.Errorf("error in box %d: %w", box, err)
		} else if changed {
			return changed, err
		}
	}

	return false, nil
}

func HiddenSingle(grid *sudoku.Grid) (bool, error) {
	if changed, err := findHiddenSingleInRows(grid); err != nil {
		return false, err
	} else if changed {
		return true, nil
	}

	if changed, err := findHiddenSingleInColumns(grid); err != nil {
		return false, err
	} else if changed {
		return true, nil
	}

	if changed, err := findHiddenSingleInBoxes(grid); err != nil {
		return false, err
	} else if changed {
		return true, nil
	}

	return false, nil
}
