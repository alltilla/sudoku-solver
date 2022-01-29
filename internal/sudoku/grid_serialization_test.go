package sudoku

import (
	"testing"

	. "github.com/alltilla/sudoku-solver/internal/test_utils"
)

func setCell(grid *Grid, row int, column int, value int, pencil_marks []int) {
	cell, err := grid.GetCell(row, column)
	if err != nil {
		panic(err.Error())
	}

	if value != Empty {
		err = cell.SetValue(value)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	if len(pencil_marks) > 0 {
		err := cell.RemovePencilMarks([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		if err != nil {
			panic(err.Error())
		}

		err = cell.AddPencilMarks(pencil_marks)
		if err != nil {
			panic(err.Error())
		}
		return
	}
}

func TestLoadPrettyString(t *testing.T) {
	expected_grid := NewGrid()
	setCell(expected_grid, 1, 1, 1, []int{})
	setCell(expected_grid, 1, 2, Empty, []int{1, 2})
	setCell(expected_grid, 1, 3, Empty, []int{1, 3})
	setCell(expected_grid, 1, 4, Empty, []int{1, 4})
	setCell(expected_grid, 1, 5, Empty, []int{1, 5})
	setCell(expected_grid, 1, 6, Empty, []int{1, 6})
	setCell(expected_grid, 1, 7, Empty, []int{1, 7})
	setCell(expected_grid, 1, 8, Empty, []int{1, 8})
	setCell(expected_grid, 1, 9, Empty, []int{1, 9})
	setCell(expected_grid, 2, 1, Empty, []int{2, 3})
	setCell(expected_grid, 2, 2, 2, []int{})
	setCell(expected_grid, 2, 3, Empty, []int{2, 4})
	setCell(expected_grid, 2, 4, Empty, []int{2, 5})
	setCell(expected_grid, 2, 5, Empty, []int{2, 6})
	setCell(expected_grid, 2, 6, Empty, []int{2, 7})
	setCell(expected_grid, 2, 7, Empty, []int{2, 8})
	setCell(expected_grid, 2, 8, Empty, []int{2, 9})
	setCell(expected_grid, 2, 9, Empty, []int{3, 4})
	setCell(expected_grid, 3, 1, Empty, []int{3, 5})
	setCell(expected_grid, 3, 2, Empty, []int{3, 6})
	setCell(expected_grid, 3, 3, 3, []int{})
	setCell(expected_grid, 3, 4, Empty, []int{3, 7})
	setCell(expected_grid, 3, 5, Empty, []int{3, 8})
	setCell(expected_grid, 3, 6, Empty, []int{3, 9})
	setCell(expected_grid, 3, 7, Empty, []int{4, 5})
	setCell(expected_grid, 3, 8, Empty, []int{4, 6})
	setCell(expected_grid, 3, 9, Empty, []int{4, 7})
	setCell(expected_grid, 4, 1, Empty, []int{4, 8})
	setCell(expected_grid, 4, 2, Empty, []int{4, 9})
	setCell(expected_grid, 4, 3, Empty, []int{5, 6})
	setCell(expected_grid, 4, 4, 4, []int{})
	setCell(expected_grid, 4, 5, Empty, []int{5, 7})
	setCell(expected_grid, 4, 6, Empty, []int{5, 8})
	setCell(expected_grid, 4, 7, Empty, []int{5, 9})
	setCell(expected_grid, 4, 8, Empty, []int{6, 7})
	setCell(expected_grid, 4, 9, Empty, []int{6, 8})
	setCell(expected_grid, 5, 1, Empty, []int{6, 9})
	setCell(expected_grid, 5, 2, Empty, []int{7, 8})
	setCell(expected_grid, 5, 3, Empty, []int{7, 9})
	setCell(expected_grid, 5, 4, Empty, []int{8, 9})
	setCell(expected_grid, 5, 5, 5, []int{})
	setCell(expected_grid, 5, 6, Empty, []int{1, 2, 3})
	setCell(expected_grid, 5, 7, Empty, []int{1, 2, 4})
	setCell(expected_grid, 5, 8, Empty, []int{1, 2, 5})
	setCell(expected_grid, 5, 9, Empty, []int{1, 2, 6})
	setCell(expected_grid, 6, 1, Empty, []int{1, 2, 7})
	setCell(expected_grid, 6, 2, Empty, []int{1, 2, 8})
	setCell(expected_grid, 6, 3, Empty, []int{1, 2, 9})
	setCell(expected_grid, 6, 4, Empty, []int{1, 3, 4})
	setCell(expected_grid, 6, 5, Empty, []int{1, 3, 5})
	setCell(expected_grid, 6, 6, 6, []int{})
	setCell(expected_grid, 6, 7, Empty, []int{1, 3, 6})
	setCell(expected_grid, 6, 8, Empty, []int{1, 3, 7})
	setCell(expected_grid, 6, 9, Empty, []int{1, 3, 8})
	setCell(expected_grid, 7, 1, Empty, []int{1, 3, 9})
	setCell(expected_grid, 7, 2, Empty, []int{1, 4, 5})
	setCell(expected_grid, 7, 3, Empty, []int{1, 4, 6})
	setCell(expected_grid, 7, 4, Empty, []int{1, 4, 7})
	setCell(expected_grid, 7, 5, Empty, []int{1, 4, 8})
	setCell(expected_grid, 7, 6, Empty, []int{1, 4, 9})
	setCell(expected_grid, 7, 7, 7, []int{})
	setCell(expected_grid, 7, 8, Empty, []int{1, 5, 6})
	setCell(expected_grid, 7, 9, Empty, []int{1, 5, 7})
	setCell(expected_grid, 8, 1, Empty, []int{1, 5, 8})
	setCell(expected_grid, 8, 2, Empty, []int{1, 5, 9})
	setCell(expected_grid, 8, 3, Empty, []int{1, 6, 7})
	setCell(expected_grid, 8, 4, Empty, []int{1, 6, 8})
	setCell(expected_grid, 8, 5, Empty, []int{1, 6, 9})
	setCell(expected_grid, 8, 6, Empty, []int{1, 7, 8})
	setCell(expected_grid, 8, 7, Empty, []int{1, 7, 9})
	setCell(expected_grid, 8, 8, 8, []int{})
	setCell(expected_grid, 8, 9, Empty, []int{1, 8, 9})
	setCell(expected_grid, 9, 1, Empty, []int{2, 3, 4})
	setCell(expected_grid, 9, 2, Empty, []int{2, 3, 5})
	setCell(expected_grid, 9, 3, Empty, []int{2, 3, 6})
	setCell(expected_grid, 9, 4, Empty, []int{2, 3, 7})
	setCell(expected_grid, 9, 5, Empty, []int{2, 3, 8})
	setCell(expected_grid, 9, 6, Empty, []int{2, 3, 9})
	setCell(expected_grid, 9, 7, Empty, []int{2, 4, 5})
	setCell(expected_grid, 9, 8, Empty, []int{2, 4, 6})
	setCell(expected_grid, 9, 9, 9, []int{})

	grid := NewGrid()
	AssertNoError(t, grid.LoadPrettyString(TEST_GRID_PRETTY_STRING))

	if !grid.Equals(expected_grid) {
		t.Errorf("Grids do not match")
	}
}
