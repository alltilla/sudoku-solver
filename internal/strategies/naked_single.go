package strategies

import (
	"fmt"

	"github.com/alltilla/sudoku-solver/internal/sudoku"
)

func NakedSingle(grid *sudoku.Grid) (bool, error) {
	for _, cell := range grid.GetAllCells() {
		if cell.GetValue() != sudoku.Empty {
			continue
		}

		pencil_marks := cell.GetPencilMarks()

		if len(pencil_marks) == 1 {
			cell.SetValue(pencil_marks[0])
			return true, nil
		}

		if len(pencil_marks) == 0 {
			return false, fmt.Errorf("no possible value for cell (%d, %d)", cell.GetRowId(), cell.GetColumnId())
		}
	}

	return false, nil
}
