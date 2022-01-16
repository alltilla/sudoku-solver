package sudoku

import (
	"fmt"
	"strconv"
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

func TestLoadValuesFromString(t *testing.T) {
	input_string := `   26 7 1
68  7  9 
19   45  
82 1   4 
  46 29  
 5   3 28
  93   74
 4  5  36
7 3 18   
`
	expected_cells := []struct {
		row    int
		column int
		value  int
	}{
		{1, 1, Empty}, {1, 2, Empty}, {1, 3, Empty}, {1, 4, 2}, {1, 5, 6}, {1, 6, Empty}, {1, 7, 7}, {1, 8, Empty}, {1, 9, 1},
		{2, 1, 6}, {2, 2, 8}, {2, 3, Empty}, {2, 4, Empty}, {2, 5, 7}, {2, 6, Empty}, {2, 7, Empty}, {2, 8, 9}, {2, 9, Empty},
		{3, 1, 1}, {3, 2, 9}, {3, 3, Empty}, {3, 4, Empty}, {3, 5, Empty}, {3, 6, 4}, {3, 7, 5}, {3, 8, Empty}, {3, 9, Empty},
		{4, 1, 8}, {4, 2, 2}, {4, 3, Empty}, {4, 4, 1}, {4, 5, Empty}, {4, 6, Empty}, {4, 7, Empty}, {4, 8, 4}, {4, 9, Empty},
		{5, 1, Empty}, {5, 2, Empty}, {5, 3, 4}, {5, 4, 6}, {5, 5, Empty}, {5, 6, 2}, {5, 7, 9}, {5, 8, Empty}, {5, 9, Empty},
		{6, 1, Empty}, {6, 2, 5}, {6, 3, Empty}, {6, 4, Empty}, {6, 5, Empty}, {6, 6, 3}, {6, 7, Empty}, {6, 8, 2}, {6, 9, 8},
		{7, 1, Empty}, {7, 2, Empty}, {7, 3, 9}, {7, 4, 3}, {7, 5, Empty}, {7, 6, Empty}, {7, 7, Empty}, {7, 8, 7}, {7, 9, 4},
		{8, 1, Empty}, {8, 2, 4}, {8, 3, Empty}, {8, 4, Empty}, {8, 5, 5}, {8, 6, Empty}, {8, 7, Empty}, {8, 8, 3}, {8, 9, 6},
		{9, 1, 7}, {9, 2, Empty}, {9, 3, 3}, {9, 4, Empty}, {9, 5, 1}, {9, 6, 8}, {9, 7, Empty}, {9, 8, Empty}, {9, 9, Empty},
	}

	grid := NewGrid()
	AssertNoError(t, grid.LoadValuesFromString(input_string))

	for _, expected_cell := range expected_cells {
		cell, err := grid.GetCell(expected_cell.row, expected_cell.column)
		AssertNoError(t, err)
		AssertValue(t, cell.GetValue(), expected_cell.value)
	}
}

func TestGetCellsInRow(t *testing.T) {
	grid_string := `   26 7 1
68  7  9 
19   45  
82 1   4 
  46 29  
 5   3 28
  93   74
 4  5  36
7 3 18   
`

	grid := NewGrid()
	grid.LoadValuesFromString(grid_string)

	cells, err := grid.GetCellsInRow(5)
	AssertNoError(t, err)

	expected_values := [9]int{Empty, Empty, 4, 6, Empty, 2, 9, Empty, Empty}
	for i, cell := range cells {
		AssertValue(t, cell.GetValue(), expected_values[i])
	}
}

func TestGetCellsInRowInvalid(t *testing.T) {
	rows_to_get := []int{0, 10}
	for _, row_to_get := range rows_to_get {
		t.Run(strconv.Itoa(row_to_get), func(t *testing.T) {
			grid := NewGrid()

			cells, err := grid.GetCellsInRow(row_to_get)
			AssertError(t, err)

			for _, cell := range cells {
				if cell != nil {
					t.Errorf("unexpected cells: %v", cells)
				}
			}
		})
	}
}

func TestGetCellsInColumn(t *testing.T) {
	grid_string := `   26 7 1
68  7  9 
19   45  
82 1   4 
  46 29  
 5   3 28
  93   74
 4  5  36
7 3 18   
`

	grid := NewGrid()
	grid.LoadValuesFromString(grid_string)

	cells, err := grid.GetCellsInColumn(5)
	AssertNoError(t, err)

	expected_values := [9]int{6, 7, Empty, Empty, Empty, Empty, Empty, 5, 1}
	for i, cell := range cells {
		AssertValue(t, cell.GetValue(), expected_values[i])
	}
}

func TestGetColumnInvalid(t *testing.T) {
	columns_to_get := []int{0, 10}
	for _, column_to_get := range columns_to_get {
		t.Run(strconv.Itoa(column_to_get), func(t *testing.T) {
			grid := NewGrid()

			cells, err := grid.GetCellsInColumn(column_to_get)
			AssertError(t, err)

			for _, cell := range cells {
				if cell != nil {
					t.Errorf("unexpected cells: %v", cells)
				}
			}
		})
	}
}
