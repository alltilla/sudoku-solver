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

func TestGridLoadValueFromString(t *testing.T) {
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

func TestGridLoadPencilMarksFromString(t *testing.T) {
	input_string := `1     | 1 2   | 1 3   | 1 4   | 1 5   | 1 6   | 1 7   | 1 8     | 1 9 
                     1 2 3 | 2     | 2 3   | 2 4   | 2 5   | 2 6   | 2 7   | 2 8     | 2 9
                     1 3 4 | 2 3 4 | 3     | 3 4   | 3 5   | 3 6   | 3 7   | 3 8     | 3 9
                     1 4 5 | 2 4 5 | 3 4 5 | 4     | 4 5   | 4 6   | 4 7   | 4 8     | 4 9
                     1 5 6 | 2 5 6 | 3 5 6 | 4 5 6 | 5     | 5 6   | 5 7   | 5 8     | 5 9
                     1 6 7 | 2 6 7 | 3 6 7 | 4 6 7 | 5 6 7 | 6     | 6 7   | 6 8     | 6 9
                     1 7 8 | 2 7 8 | 3 7 8 | 4 7 8 | 5 7 8 | 6 7 8 | 7     | 7 8     | 7 9
                     1 8 9 | 2 8 9 | 3 8 9 | 4 8 9 | 5 8 9 | 6 8 9 | 7 8 9 | 8       | 8 9
                     1 2 9 | 2 3 9 | 3 4 9 | 4 5 9 | 5 6 9 | 6 7 9 | 1 7 9 | 1 2 8 9 |    
`
	expected_cells := []struct {
		row          int
		column       int
		pencil_marks []int
	}{
		{1, 1, []int{1}}, {1, 2, []int{1, 2}}, {1, 3, []int{1, 3}}, {1, 4, []int{1, 4}}, {1, 5, []int{1, 5}}, {1, 6, []int{1, 6}}, {1, 7, []int{1, 7}}, {1, 8, []int{1, 8}}, {1, 9, []int{1, 9}},
		{2, 1, []int{1, 2, 3}}, {2, 2, []int{2}}, {2, 3, []int{2, 3}}, {2, 4, []int{2, 4}}, {2, 5, []int{2, 5}}, {2, 6, []int{2, 6}}, {2, 7, []int{2, 7}}, {2, 8, []int{2, 8}}, {2, 9, []int{2, 9}},
		{3, 1, []int{1, 3, 4}}, {3, 2, []int{2, 3, 4}}, {3, 3, []int{3}}, {3, 4, []int{3, 4}}, {3, 5, []int{3, 5}}, {3, 6, []int{3, 6}}, {3, 7, []int{3, 7}}, {3, 8, []int{3, 8}}, {3, 9, []int{3, 9}},
		{4, 1, []int{1, 4, 5}}, {4, 2, []int{2, 4, 5}}, {4, 3, []int{3, 4, 5}}, {4, 4, []int{4}}, {4, 5, []int{4, 5}}, {4, 6, []int{4, 6}}, {4, 7, []int{4, 7}}, {4, 8, []int{4, 8}}, {4, 9, []int{4, 9}},
		{5, 1, []int{1, 5, 6}}, {5, 2, []int{2, 5, 6}}, {5, 3, []int{3, 5, 6}}, {5, 4, []int{4, 5, 6}}, {5, 5, []int{5}}, {5, 6, []int{5, 6}}, {5, 7, []int{5, 7}}, {5, 8, []int{5, 8}}, {5, 9, []int{5, 9}},
		{6, 1, []int{1, 6, 7}}, {6, 2, []int{2, 6, 7}}, {6, 3, []int{3, 6, 7}}, {6, 4, []int{4, 6, 7}}, {6, 5, []int{5, 6, 7}}, {6, 6, []int{6}}, {6, 7, []int{6, 7}}, {6, 8, []int{6, 8}}, {6, 9, []int{6, 9}},
		{7, 1, []int{1, 7, 8}}, {7, 2, []int{2, 7, 8}}, {7, 3, []int{3, 7, 8}}, {7, 4, []int{4, 7, 8}}, {7, 5, []int{5, 7, 8}}, {7, 6, []int{6, 7, 8}}, {7, 7, []int{7}}, {7, 8, []int{7, 8}}, {7, 9, []int{7, 9}},
		{8, 1, []int{1, 8, 9}}, {8, 2, []int{2, 8, 9}}, {8, 3, []int{3, 8, 9}}, {8, 4, []int{4, 8, 9}}, {8, 5, []int{5, 8, 9}}, {8, 6, []int{6, 8, 9}}, {8, 7, []int{7, 8, 9}}, {8, 8, []int{8}}, {8, 9, []int{8, 9}},
		{9, 1, []int{1, 2, 9}}, {9, 2, []int{2, 3, 9}}, {9, 3, []int{3, 4, 9}}, {9, 4, []int{4, 5, 9}}, {9, 5, []int{5, 6, 9}}, {9, 6, []int{6, 7, 9}}, {9, 7, []int{1, 7, 9}}, {9, 8, []int{1, 2, 8, 9}}, {9, 9, []int{}},
	}

	grid := NewGrid()
	AssertNoError(t, grid.LoadPencilMarksFromString(input_string))

	for _, expected_cell := range expected_cells {
		cell, err := grid.GetCell(expected_cell.row, expected_cell.column)
		AssertNoError(t, err)
		assertPencilMarks(t, cell.GetPencilMarks(), expected_cell.pencil_marks)
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

func TestGetCellsInBox(t *testing.T) {
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

	cells, err := grid.GetCellsInBox(5)
	AssertNoError(t, err)

	expected_values := [9]int{1, Empty, Empty, 6, Empty, 2, Empty, Empty, 3}
	for i, cell := range cells {
		AssertValue(t, cell.GetValue(), expected_values[i])
	}
}

func TestGetBoxInvalid(t *testing.T) {
	boxes_to_get := []int{0, 10}
	for _, box_to_get := range boxes_to_get {
		t.Run(strconv.Itoa(box_to_get), func(t *testing.T) {
			grid := NewGrid()

			box, err := grid.GetCellsInBox(box_to_get)
			AssertError(t, err)

			for _, cell := range box {
				if cell != nil {
					t.Errorf("unexpected box: %v", box)
				}
			}
		})
	}
}

func createCell(row int, column int, value int) *Cell {
	cell, err := NewCell(row, column)
	if err != nil {
		panic("failed to create cell")
	}

	err = cell.SetValue(value)
	if err != nil {
		panic("failed to set value")
	}

	return cell
}

func TestGetAllCells(t *testing.T) {
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

	expected_cells := [81]*Cell{
		createCell(1, 1, Empty), createCell(1, 2, Empty), createCell(1, 3, Empty), createCell(1, 4, 2), createCell(1, 5, 6), createCell(1, 6, Empty), createCell(1, 7, 7), createCell(1, 8, Empty), createCell(1, 9, 1),
		createCell(2, 1, 6), createCell(2, 2, 8), createCell(2, 3, Empty), createCell(2, 4, Empty), createCell(2, 5, 7), createCell(2, 6, Empty), createCell(2, 7, Empty), createCell(2, 8, 9), createCell(2, 9, Empty),
		createCell(3, 1, 1), createCell(3, 2, 9), createCell(3, 3, Empty), createCell(3, 4, Empty), createCell(3, 5, Empty), createCell(3, 6, 4), createCell(3, 7, 5), createCell(3, 8, Empty), createCell(3, 9, Empty),
		createCell(4, 1, 8), createCell(4, 2, 2), createCell(4, 3, Empty), createCell(4, 4, 1), createCell(4, 5, Empty), createCell(4, 6, Empty), createCell(4, 7, Empty), createCell(4, 8, 4), createCell(4, 9, Empty),
		createCell(5, 1, Empty), createCell(5, 2, Empty), createCell(5, 3, 4), createCell(5, 4, 6), createCell(5, 5, Empty), createCell(5, 6, 2), createCell(5, 7, 9), createCell(5, 8, Empty), createCell(5, 9, Empty),
		createCell(6, 1, Empty), createCell(6, 2, 5), createCell(6, 3, Empty), createCell(6, 4, Empty), createCell(6, 5, Empty), createCell(6, 6, 3), createCell(6, 7, Empty), createCell(6, 8, 2), createCell(6, 9, 8),
		createCell(7, 1, Empty), createCell(7, 2, Empty), createCell(7, 3, 9), createCell(7, 4, 3), createCell(7, 5, Empty), createCell(7, 6, Empty), createCell(7, 7, Empty), createCell(7, 8, 7), createCell(7, 9, 4),
		createCell(8, 1, Empty), createCell(8, 2, 4), createCell(8, 3, Empty), createCell(8, 4, Empty), createCell(8, 5, 5), createCell(8, 6, Empty), createCell(8, 7, Empty), createCell(8, 8, 3), createCell(8, 9, 6),
		createCell(9, 1, 7), createCell(9, 2, Empty), createCell(9, 3, 3), createCell(9, 4, Empty), createCell(9, 5, 1), createCell(9, 6, 8), createCell(9, 7, Empty), createCell(9, 8, Empty), createCell(9, 9, Empty),
	}

	cells := grid.GetAllCells()

	for i, cell := range cells {
		expected_cell := expected_cells[i]
		if !cell.Equals(expected_cell) {
			t.Errorf("unexpected cell. expected: (row: %d, column: %d, value: %d), actual: (row: %d, column: %d, value: %d)",
				expected_cell.GetRowId(), expected_cell.GetColumnId(), expected_cell.GetValue(),
				cell.GetRowId(), cell.GetColumnId(), cell.GetValue(),
			)
		}
	}
}
