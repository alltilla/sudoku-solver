package sudoku

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/alltilla/sudoku-solver/internal/test_utils"
)

const TEST_GRID_PRETTY_STRING = `
##=======================##=======================##=======================##
||       | 1 2   | 1   3 || 1     | 1     | 1     || 1     | 1     | 1     ||
||  (1)  |       |       || 4     |   5   |     6 ||       |       |       ||
||       |       |       ||       |       |       || 7     |   8   |     9 ||
||-------+-------+-------||-------+-------+-------||-------+-------+-------||
||   2 3 |       |   2   ||   2   |   2   |   2   ||   2   |   2   |     3 ||
||       |  (2)  | 4     ||   5   |     6 |       ||       |       | 4     ||
||       |       |       ||       |       | 7     ||   8   |     9 |       ||
||-------+-------+-------||-------+-------+-------||-------+-------+-------||
||     3 |     3 |       ||     3 |     3 |     3 ||       |       |       ||
||   5   |     6 |  (3)  ||       |       |       || 4 5   | 4   6 | 4     ||
||       |       |       || 7     |   8   |     9 ||       |       | 7     ||
##=======================##=======================##=======================##
||       |       |       ||       |       |       ||       |       |       ||
|| 4     | 4     |   5 6 ||  (4)  |   5   |   5   ||   5   |     6 |     6 ||
||   8   |     9 |       ||       | 7     |   8   ||     9 | 7     |   8   ||
||-------+-------+-------||-------+-------+-------||-------+-------+-------||
||       |       |       ||       |       | 1 2 3 || 1 2   | 1 2   | 1 2   ||
||     6 |       |       ||       |  (5)  |       || 4     |   5   |     6 ||
||     9 | 7 8   | 7   9 ||   8 9 |       |       ||       |       |       ||
||-------+-------+-------||-------+-------+-------||-------+-------+-------||
|| 1 2   | 1 2   | 1 2   || 1   3 | 1   3 |       || 1   3 | 1   3 | 1   3 ||
||       |       |       || 4     |   5   |  (6)  ||     6 |       |       ||
|| 7     |   8   |     9 ||       |       |       ||       | 7     |   8   ||
##=======================##=======================##=======================##
|| 1   3 | 1     | 1     || 1     | 1     | 1     ||       | 1     | 1     ||
||       | 4 5   | 4   6 || 4     | 4     | 4     ||  (7)  |   5 6 |   5   ||
||     9 |       |       || 7     |   8   |     9 ||       |       | 7     ||
||-------+-------+-------||-------+-------+-------||-------+-------+-------||
|| 1     | 1     | 1     || 1     | 1     | 1     || 1     |       | 1     ||
||   5   |   5   |     6 ||     6 |     6 |       ||       |  (8)  |       ||
||   8   |     9 | 7     ||   8   |     9 | 7 8   || 7   9 |       |   8 9 ||
||-------+-------+-------||-------+-------+-------||-------+-------+-------||
||   2 3 |   2 3 |   2 3 ||   2 3 |   2 3 |   2 3 ||   2   |   2   |       ||
|| 4     |   5   |     6 ||       |       |       || 4 5   | 4   6 |  (9)  ||
||       |       |       || 7     |   8   |     9 ||       |       |       ||
##=======================##=======================##=======================##
`

func createCell(row int, column int, value int, pencil_marks []int) *Cell {
	cell, err := NewCell(row, column)
	if err != nil {
		panic(err.Error())
	}

	if value != Empty {
		err = cell.SetValue(value)
		if err != nil {
			panic(err.Error())
		}
		return cell
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
		return cell
	}

	return cell
}

func TestGetCell(t *testing.T) {
	for row := 1; row <= 9; row++ {
		for column := 1; column <= 9; column++ {
			t.Run(fmt.Sprintf("%d,%d", row, column), func(t *testing.T) {
				grid := NewGrid()

				cell, err := grid.GetCell(row, column)
				AssertNoError(t, err)

				expected_cell := createCell(row, column, Empty, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
				AssertCellEquals(t, cell, expected_cell)
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

func TestGetCellsInRow(t *testing.T) {
	grid := NewGrid()
	AssertNoError(t, grid.LoadPrettyString(TEST_GRID_PRETTY_STRING))

	cells, err := grid.GetCellsInRow(5)
	AssertNoError(t, err)

	expected_cells := [9]*Cell{
		createCell(5, 1, Empty, []int{6, 9}),
		createCell(5, 2, Empty, []int{7, 8}),
		createCell(5, 3, Empty, []int{7, 9}),
		createCell(5, 4, Empty, []int{8, 9}),
		createCell(5, 5, 5, []int{}),
		createCell(5, 6, Empty, []int{1, 2, 3}),
		createCell(5, 7, Empty, []int{1, 2, 4}),
		createCell(5, 8, Empty, []int{1, 2, 5}),
		createCell(5, 9, Empty, []int{1, 2, 6}),
	}

	for i, cell := range cells {
		AssertCellEquals(t, cell, expected_cells[i])
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
	grid := NewGrid()
	AssertNoError(t, grid.LoadPrettyString(TEST_GRID_PRETTY_STRING))

	cells, err := grid.GetCellsInColumn(5)
	AssertNoError(t, err)

	expected_cells := [9]*Cell{
		createCell(1, 5, Empty, []int{1, 5}),
		createCell(2, 5, Empty, []int{2, 6}),
		createCell(3, 5, Empty, []int{3, 8}),
		createCell(4, 5, Empty, []int{5, 7}),
		createCell(5, 5, 5, []int{}),
		createCell(6, 5, Empty, []int{1, 3, 5}),
		createCell(7, 5, Empty, []int{1, 4, 8}),
		createCell(8, 5, Empty, []int{1, 6, 9}),
		createCell(9, 5, Empty, []int{2, 3, 8}),
	}

	for i, cell := range cells {
		AssertCellEquals(t, cell, expected_cells[i])
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
	grid := NewGrid()
	AssertNoError(t, grid.LoadPrettyString(TEST_GRID_PRETTY_STRING))

	cells, err := grid.GetCellsInBox(5)
	AssertNoError(t, err)

	expected_cells := [9]*Cell{
		createCell(4, 4, 4, []int{}),
		createCell(4, 5, Empty, []int{5, 7}),
		createCell(4, 6, Empty, []int{5, 8}),
		createCell(5, 4, Empty, []int{8, 9}),
		createCell(5, 5, 5, []int{}),
		createCell(5, 6, Empty, []int{1, 2, 3}),
		createCell(6, 4, Empty, []int{1, 3, 4}),
		createCell(6, 5, Empty, []int{1, 3, 5}),
		createCell(6, 6, 6, []int{}),
	}

	for i, cell := range cells {
		AssertCellEquals(t, cell, expected_cells[i])
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

func TestGetSets(t *testing.T) {
	grid := NewGrid()
	grid.LoadPrettyString(TEST_GRID_PRETTY_STRING)

	expected_sets := [27]*Set{
		{"row", 1, [9]*Cell{createCell(1, 1, 1, []int{}), createCell(1, 2, Empty, []int{1, 2}), createCell(1, 3, Empty, []int{1, 3}), createCell(1, 4, Empty, []int{1, 4}), createCell(1, 5, Empty, []int{1, 5}), createCell(1, 6, Empty, []int{1, 6}), createCell(1, 7, Empty, []int{1, 7}), createCell(1, 8, Empty, []int{1, 8}), createCell(1, 9, Empty, []int{1, 9})}},
		{"row", 2, [9]*Cell{createCell(2, 1, Empty, []int{2, 3}), createCell(2, 2, 2, []int{}), createCell(2, 3, Empty, []int{2, 4}), createCell(2, 4, Empty, []int{2, 5}), createCell(2, 5, Empty, []int{2, 6}), createCell(2, 6, Empty, []int{2, 7}), createCell(2, 7, Empty, []int{2, 8}), createCell(2, 8, Empty, []int{2, 9}), createCell(2, 9, Empty, []int{3, 4})}},
		{"row", 3, [9]*Cell{createCell(3, 1, Empty, []int{3, 5}), createCell(3, 2, Empty, []int{3, 6}), createCell(3, 3, 3, []int{}), createCell(3, 4, Empty, []int{3, 7}), createCell(3, 5, Empty, []int{3, 8}), createCell(3, 6, Empty, []int{3, 9}), createCell(3, 7, Empty, []int{4, 5}), createCell(3, 8, Empty, []int{4, 6}), createCell(3, 9, Empty, []int{4, 7})}},
		{"row", 4, [9]*Cell{createCell(4, 1, Empty, []int{4, 8}), createCell(4, 2, Empty, []int{4, 9}), createCell(4, 3, Empty, []int{5, 6}), createCell(4, 4, 4, []int{}), createCell(4, 5, Empty, []int{5, 7}), createCell(4, 6, Empty, []int{5, 8}), createCell(4, 7, Empty, []int{5, 9}), createCell(4, 8, Empty, []int{6, 7}), createCell(4, 9, Empty, []int{6, 8})}},
		{"row", 5, [9]*Cell{createCell(5, 1, Empty, []int{6, 9}), createCell(5, 2, Empty, []int{7, 8}), createCell(5, 3, Empty, []int{7, 9}), createCell(5, 4, Empty, []int{8, 9}), createCell(5, 5, 5, []int{}), createCell(5, 6, Empty, []int{1, 2, 3}), createCell(5, 7, Empty, []int{1, 2, 4}), createCell(5, 8, Empty, []int{1, 2, 5}), createCell(5, 9, Empty, []int{1, 2, 6})}},
		{"row", 6, [9]*Cell{createCell(6, 1, Empty, []int{1, 2, 7}), createCell(6, 2, Empty, []int{1, 2, 8}), createCell(6, 3, Empty, []int{1, 2, 9}), createCell(6, 4, Empty, []int{1, 3, 4}), createCell(6, 5, Empty, []int{1, 3, 5}), createCell(6, 6, 6, []int{}), createCell(6, 7, Empty, []int{1, 3, 6}), createCell(6, 8, Empty, []int{1, 3, 7}), createCell(6, 9, Empty, []int{1, 3, 8})}},
		{"row", 7, [9]*Cell{createCell(7, 1, Empty, []int{1, 3, 9}), createCell(7, 2, Empty, []int{1, 4, 5}), createCell(7, 3, Empty, []int{1, 4, 6}), createCell(7, 4, Empty, []int{1, 4, 7}), createCell(7, 5, Empty, []int{1, 4, 8}), createCell(7, 6, Empty, []int{1, 4, 9}), createCell(7, 7, 7, []int{}), createCell(7, 8, Empty, []int{1, 5, 6}), createCell(7, 9, Empty, []int{1, 5, 7})}},
		{"row", 8, [9]*Cell{createCell(8, 1, Empty, []int{1, 5, 8}), createCell(8, 2, Empty, []int{1, 5, 9}), createCell(8, 3, Empty, []int{1, 6, 7}), createCell(8, 4, Empty, []int{1, 6, 8}), createCell(8, 5, Empty, []int{1, 6, 9}), createCell(8, 6, Empty, []int{1, 7, 8}), createCell(8, 7, Empty, []int{1, 7, 9}), createCell(8, 8, 8, []int{}), createCell(8, 9, Empty, []int{1, 8, 9})}},
		{"row", 9, [9]*Cell{createCell(9, 1, Empty, []int{2, 3, 4}), createCell(9, 2, Empty, []int{2, 3, 5}), createCell(9, 3, Empty, []int{2, 3, 6}), createCell(9, 4, Empty, []int{2, 3, 7}), createCell(9, 5, Empty, []int{2, 3, 8}), createCell(9, 6, Empty, []int{2, 3, 9}), createCell(9, 7, Empty, []int{2, 4, 5}), createCell(9, 8, Empty, []int{2, 4, 6}), createCell(9, 9, 9, []int{})}},
		{"column", 1, [9]*Cell{createCell(1, 1, 1, []int{}), createCell(2, 1, Empty, []int{2, 3}), createCell(3, 1, Empty, []int{3, 5}), createCell(4, 1, Empty, []int{4, 8}), createCell(5, 1, Empty, []int{6, 9}), createCell(6, 1, Empty, []int{1, 2, 7}), createCell(7, 1, Empty, []int{1, 3, 9}), createCell(8, 1, Empty, []int{1, 5, 8}), createCell(9, 1, Empty, []int{2, 3, 4})}},
		{"column", 2, [9]*Cell{createCell(1, 2, Empty, []int{1, 2}), createCell(2, 2, 2, []int{}), createCell(3, 2, Empty, []int{3, 6}), createCell(4, 2, Empty, []int{4, 9}), createCell(5, 2, Empty, []int{7, 8}), createCell(6, 2, Empty, []int{1, 2, 8}), createCell(7, 2, Empty, []int{1, 4, 5}), createCell(8, 2, Empty, []int{1, 5, 9}), createCell(9, 2, Empty, []int{2, 3, 5})}},
		{"column", 3, [9]*Cell{createCell(1, 3, Empty, []int{1, 3}), createCell(2, 3, Empty, []int{2, 4}), createCell(3, 3, 3, []int{}), createCell(4, 3, Empty, []int{5, 6}), createCell(5, 3, Empty, []int{7, 9}), createCell(6, 3, Empty, []int{1, 2, 9}), createCell(7, 3, Empty, []int{1, 4, 6}), createCell(8, 3, Empty, []int{1, 6, 7}), createCell(9, 3, Empty, []int{2, 3, 6})}},
		{"column", 4, [9]*Cell{createCell(1, 4, Empty, []int{1, 4}), createCell(2, 4, Empty, []int{2, 5}), createCell(3, 4, Empty, []int{3, 7}), createCell(4, 4, 4, []int{}), createCell(5, 4, Empty, []int{8, 9}), createCell(6, 4, Empty, []int{1, 3, 4}), createCell(7, 4, Empty, []int{1, 4, 7}), createCell(8, 4, Empty, []int{1, 6, 8}), createCell(9, 4, Empty, []int{2, 3, 7})}},
		{"column", 5, [9]*Cell{createCell(1, 5, Empty, []int{1, 5}), createCell(2, 5, Empty, []int{2, 6}), createCell(3, 5, Empty, []int{3, 8}), createCell(4, 5, Empty, []int{5, 7}), createCell(5, 5, 5, []int{}), createCell(6, 5, Empty, []int{1, 3, 5}), createCell(7, 5, Empty, []int{1, 4, 8}), createCell(8, 5, Empty, []int{1, 6, 9}), createCell(9, 5, Empty, []int{2, 3, 8})}},
		{"column", 6, [9]*Cell{createCell(1, 6, Empty, []int{1, 6}), createCell(2, 6, Empty, []int{2, 7}), createCell(3, 6, Empty, []int{3, 9}), createCell(4, 6, Empty, []int{5, 8}), createCell(5, 6, Empty, []int{1, 2, 3}), createCell(6, 6, 6, []int{}), createCell(7, 6, Empty, []int{1, 4, 9}), createCell(8, 6, Empty, []int{1, 7, 8}), createCell(9, 6, Empty, []int{2, 3, 9})}},
		{"column", 7, [9]*Cell{createCell(1, 7, Empty, []int{1, 7}), createCell(2, 7, Empty, []int{2, 8}), createCell(3, 7, Empty, []int{4, 5}), createCell(4, 7, Empty, []int{5, 9}), createCell(5, 7, Empty, []int{1, 2, 4}), createCell(6, 7, Empty, []int{1, 3, 6}), createCell(7, 7, 7, []int{}), createCell(8, 7, Empty, []int{1, 7, 9}), createCell(9, 7, Empty, []int{2, 4, 5})}},
		{"column", 8, [9]*Cell{createCell(1, 8, Empty, []int{1, 8}), createCell(2, 8, Empty, []int{2, 9}), createCell(3, 8, Empty, []int{4, 6}), createCell(4, 8, Empty, []int{6, 7}), createCell(5, 8, Empty, []int{1, 2, 5}), createCell(6, 8, Empty, []int{1, 3, 7}), createCell(7, 8, Empty, []int{1, 5, 6}), createCell(8, 8, 8, []int{}), createCell(9, 8, Empty, []int{2, 4, 6})}},
		{"column", 9, [9]*Cell{createCell(1, 9, Empty, []int{1, 9}), createCell(2, 9, Empty, []int{3, 4}), createCell(3, 9, Empty, []int{4, 7}), createCell(4, 9, Empty, []int{6, 8}), createCell(5, 9, Empty, []int{1, 2, 6}), createCell(6, 9, Empty, []int{1, 3, 8}), createCell(7, 9, Empty, []int{1, 5, 7}), createCell(8, 9, Empty, []int{1, 8, 9}), createCell(9, 9, 9, []int{})}},
		{"box", 1, [9]*Cell{createCell(1, 1, 1, []int{}), createCell(1, 2, Empty, []int{1, 2}), createCell(1, 3, Empty, []int{1, 3}), createCell(2, 1, Empty, []int{2, 3}), createCell(2, 2, 2, []int{}), createCell(2, 3, Empty, []int{2, 4}), createCell(3, 1, Empty, []int{3, 5}), createCell(3, 2, Empty, []int{3, 6}), createCell(3, 3, 3, []int{})}},
		{"box", 2, [9]*Cell{createCell(1, 4, Empty, []int{1, 4}), createCell(1, 5, Empty, []int{1, 5}), createCell(1, 6, Empty, []int{1, 6}), createCell(2, 4, Empty, []int{2, 5}), createCell(2, 5, Empty, []int{2, 6}), createCell(2, 6, Empty, []int{2, 7}), createCell(3, 4, Empty, []int{3, 7}), createCell(3, 5, Empty, []int{3, 8}), createCell(3, 6, Empty, []int{3, 9})}},
		{"box", 3, [9]*Cell{createCell(1, 7, Empty, []int{1, 7}), createCell(1, 8, Empty, []int{1, 8}), createCell(1, 9, Empty, []int{1, 9}), createCell(2, 7, Empty, []int{2, 8}), createCell(2, 8, Empty, []int{2, 9}), createCell(2, 9, Empty, []int{3, 4}), createCell(3, 7, Empty, []int{4, 5}), createCell(3, 8, Empty, []int{4, 6}), createCell(3, 9, Empty, []int{4, 7})}},
		{"box", 4, [9]*Cell{createCell(4, 1, Empty, []int{4, 8}), createCell(4, 2, Empty, []int{4, 9}), createCell(4, 3, Empty, []int{5, 6}), createCell(5, 1, Empty, []int{6, 9}), createCell(5, 2, Empty, []int{7, 8}), createCell(5, 3, Empty, []int{7, 9}), createCell(6, 1, Empty, []int{1, 2, 7}), createCell(6, 2, Empty, []int{1, 2, 8}), createCell(6, 3, Empty, []int{1, 2, 9})}},
		{"box", 5, [9]*Cell{createCell(4, 4, 4, []int{}), createCell(4, 5, Empty, []int{5, 7}), createCell(4, 6, Empty, []int{5, 8}), createCell(5, 4, Empty, []int{8, 9}), createCell(5, 5, 5, []int{}), createCell(5, 6, Empty, []int{1, 2, 3}), createCell(6, 4, Empty, []int{1, 3, 4}), createCell(6, 5, Empty, []int{1, 3, 5}), createCell(6, 6, 6, []int{})}},
		{"box", 6, [9]*Cell{createCell(4, 7, Empty, []int{5, 9}), createCell(4, 8, Empty, []int{6, 7}), createCell(4, 9, Empty, []int{6, 8}), createCell(5, 7, Empty, []int{1, 2, 4}), createCell(5, 8, Empty, []int{1, 2, 5}), createCell(5, 9, Empty, []int{1, 2, 6}), createCell(6, 7, Empty, []int{1, 3, 6}), createCell(6, 8, Empty, []int{1, 3, 7}), createCell(6, 9, Empty, []int{1, 3, 8})}},
		{"box", 7, [9]*Cell{createCell(7, 1, Empty, []int{1, 3, 9}), createCell(7, 2, Empty, []int{1, 4, 5}), createCell(7, 3, Empty, []int{1, 4, 6}), createCell(8, 1, Empty, []int{1, 5, 8}), createCell(8, 2, Empty, []int{1, 5, 9}), createCell(8, 3, Empty, []int{1, 6, 7}), createCell(9, 1, Empty, []int{2, 3, 4}), createCell(9, 2, Empty, []int{2, 3, 5}), createCell(9, 3, Empty, []int{2, 3, 6})}},
		{"box", 8, [9]*Cell{createCell(7, 4, Empty, []int{1, 4, 7}), createCell(7, 5, Empty, []int{1, 4, 8}), createCell(7, 6, Empty, []int{1, 4, 9}), createCell(8, 4, Empty, []int{1, 6, 8}), createCell(8, 5, Empty, []int{1, 6, 9}), createCell(8, 6, Empty, []int{1, 7, 8}), createCell(9, 4, Empty, []int{2, 3, 7}), createCell(9, 5, Empty, []int{2, 3, 8}), createCell(9, 6, Empty, []int{2, 3, 9})}},
		{"box", 9, [9]*Cell{createCell(7, 7, 7, []int{}), createCell(7, 8, Empty, []int{1, 5, 6}), createCell(7, 9, Empty, []int{1, 5, 7}), createCell(8, 7, Empty, []int{1, 7, 9}), createCell(8, 8, 8, []int{}), createCell(8, 9, Empty, []int{1, 8, 9}), createCell(9, 7, Empty, []int{2, 4, 5}), createCell(9, 8, Empty, []int{2, 4, 6}), createCell(9, 9, 9, []int{})}},
	}

	sets := grid.GetSets()

	for i, set := range sets {
		expected_set := expected_sets[i]

		for i, cell := range set.Cells {
			expected_cell := expected_set.Cells[i]

			AssertCellEquals(t, cell, expected_cell)
		}

		if set.Index != expected_set.Index {
			t.Errorf("unexpected index. actual: %d, expected: %d", set.Index, expected_set.Index)
		}

		if set.Orientation != expected_set.Orientation {
			t.Errorf("unexpected orientation. actual: %s, expected: %s", set.Orientation, expected_set.Orientation)
		}
	}
}

func TestGetAllCells(t *testing.T) {
	grid := NewGrid()
	grid.LoadPrettyString(TEST_GRID_PRETTY_STRING)

	expected_cells := [81]*Cell{
		createCell(1, 1, 1, []int{}), createCell(1, 2, Empty, []int{1, 2}), createCell(1, 3, Empty, []int{1, 3}), createCell(1, 4, Empty, []int{1, 4}), createCell(1, 5, Empty, []int{1, 5}), createCell(1, 6, Empty, []int{1, 6}), createCell(1, 7, Empty, []int{1, 7}), createCell(1, 8, Empty, []int{1, 8}), createCell(1, 9, Empty, []int{1, 9}),
		createCell(2, 1, Empty, []int{2, 3}), createCell(2, 2, 2, []int{}), createCell(2, 3, Empty, []int{2, 4}), createCell(2, 4, Empty, []int{2, 5}), createCell(2, 5, Empty, []int{2, 6}), createCell(2, 6, Empty, []int{2, 7}), createCell(2, 7, Empty, []int{2, 8}), createCell(2, 8, Empty, []int{2, 9}), createCell(2, 9, Empty, []int{3, 4}),
		createCell(3, 1, Empty, []int{3, 5}), createCell(3, 2, Empty, []int{3, 6}), createCell(3, 3, 3, []int{}), createCell(3, 4, Empty, []int{3, 7}), createCell(3, 5, Empty, []int{3, 8}), createCell(3, 6, Empty, []int{3, 9}), createCell(3, 7, Empty, []int{4, 5}), createCell(3, 8, Empty, []int{4, 6}), createCell(3, 9, Empty, []int{4, 7}),
		createCell(4, 1, Empty, []int{4, 8}), createCell(4, 2, Empty, []int{4, 9}), createCell(4, 3, Empty, []int{5, 6}), createCell(4, 4, 4, []int{}), createCell(4, 5, Empty, []int{5, 7}), createCell(4, 6, Empty, []int{5, 8}), createCell(4, 7, Empty, []int{5, 9}), createCell(4, 8, Empty, []int{6, 7}), createCell(4, 9, Empty, []int{6, 8}),
		createCell(5, 1, Empty, []int{6, 9}), createCell(5, 2, Empty, []int{7, 8}), createCell(5, 3, Empty, []int{7, 9}), createCell(5, 4, Empty, []int{8, 9}), createCell(5, 5, 5, []int{}), createCell(5, 6, Empty, []int{1, 2, 3}), createCell(5, 7, Empty, []int{1, 2, 4}), createCell(5, 8, Empty, []int{1, 2, 5}), createCell(5, 9, Empty, []int{1, 2, 6}),
		createCell(6, 1, Empty, []int{1, 2, 7}), createCell(6, 2, Empty, []int{1, 2, 8}), createCell(6, 3, Empty, []int{1, 2, 9}), createCell(6, 4, Empty, []int{1, 3, 4}), createCell(6, 5, Empty, []int{1, 3, 5}), createCell(6, 6, 6, []int{}), createCell(6, 7, Empty, []int{1, 3, 6}), createCell(6, 8, Empty, []int{1, 3, 7}), createCell(6, 9, Empty, []int{1, 3, 8}),
		createCell(7, 1, Empty, []int{1, 3, 9}), createCell(7, 2, Empty, []int{1, 4, 5}), createCell(7, 3, Empty, []int{1, 4, 6}), createCell(7, 4, Empty, []int{1, 4, 7}), createCell(7, 5, Empty, []int{1, 4, 8}), createCell(7, 6, Empty, []int{1, 4, 9}), createCell(7, 7, 7, []int{}), createCell(7, 8, Empty, []int{1, 5, 6}), createCell(7, 9, Empty, []int{1, 5, 7}),
		createCell(8, 1, Empty, []int{1, 5, 8}), createCell(8, 2, Empty, []int{1, 5, 9}), createCell(8, 3, Empty, []int{1, 6, 7}), createCell(8, 4, Empty, []int{1, 6, 8}), createCell(8, 5, Empty, []int{1, 6, 9}), createCell(8, 6, Empty, []int{1, 7, 8}), createCell(8, 7, Empty, []int{1, 7, 9}), createCell(8, 8, 8, []int{}), createCell(8, 9, Empty, []int{1, 8, 9}),
		createCell(9, 1, Empty, []int{2, 3, 4}), createCell(9, 2, Empty, []int{2, 3, 5}), createCell(9, 3, Empty, []int{2, 3, 6}), createCell(9, 4, Empty, []int{2, 3, 7}), createCell(9, 5, Empty, []int{2, 3, 8}), createCell(9, 6, Empty, []int{2, 3, 9}), createCell(9, 7, Empty, []int{2, 4, 5}), createCell(9, 8, Empty, []int{2, 4, 6}), createCell(9, 9, 9, []int{}),
	}
	cells := grid.GetAllCells()

	for i, cell := range cells {
		AssertCellEquals(t, cell, expected_cells[i])
	}
}

func TestGridEquals(t *testing.T) {
	grid_1 := NewGrid()
	grid_1.LoadPrettyString(TEST_GRID_PRETTY_STRING)

	grid_2 := NewGrid()
	grid_2.LoadPrettyString(TEST_GRID_PRETTY_STRING)

	if !grid_1.Equals(grid_2) {
		t.Errorf("the two grids should be equal")
	}

	cell, _ := grid_2.GetCell(9, 9)
	cell.SetValue(1)

	if grid_1.Equals(grid_2) {
		t.Errorf("the two grids should not be equal")
	}
}
