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

func findHiddenSingleInSet(set *sudoku.Set) (bool, error) {
	for digit := 1; digit <= 9; digit++ {
		var hidden_single_cell_candidate *sudoku.Cell = nil
		value_found := false
		multiple_candidates := false

		for _, cell := range set.Cells {
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
			return false, fmt.Errorf("error in %s %d: no possible cell to place digit: %d", set.Orientation, set.Index, digit)
		}

		if err := hidden_single_cell_candidate.SetValue(digit); err != nil {
			panic(err.Error())
		}
		return true, nil
	}

	return false, nil
}

func HiddenSingle(grid *sudoku.Grid) (bool, error) {
	for _, set := range grid.GetSets() {
		changed, err := findHiddenSingleInSet(set)
		if err != nil || changed {
			return changed, err
		}
	}

	return false, nil
}
