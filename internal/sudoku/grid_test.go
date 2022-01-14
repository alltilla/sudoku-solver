package sudoku

import (
	"fmt"
	"testing"

	. "github.com/alltilla/sudoku-solver/internal/test_utils"
)

func TestGetCell(t *testing.T) {
	for row := 1; row <= 9; row++ {
		for column := 1; column <= 9; column++ {
			t.Run(fmt.Sprintf("%d,%d", row, column), func(t *testing.T) {
				grid := NewGrid()

				cell, err := grid.GetCell(row, column)
				AssertNoError(t, err)

				if cell == nil {
					t.Errorf("missing cell")
				}
			})
		}
	}
}

func TestGetCellInvalid(t *testing.T) {
	coordinates_to_get := [][2]int{
		{0, 1},
		{1, 0},
		{10, 1},
		{1, 10},
	}

	for _, coordinate_to_get := range coordinates_to_get {
		row := coordinate_to_get[0]
		column := coordinate_to_get[1]

		t.Run(fmt.Sprintf("%d,%d", row, column), func(t *testing.T) {
			grid := NewGrid()

			cell, err := grid.GetCell(row, column)
			AssertError(t, err)
			if cell != nil {
				t.Errorf("unexpected cell")
			}
		})
	}
}
