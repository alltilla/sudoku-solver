package sudoku

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	. "github.com/alltilla/sudoku-solver/internal/test_utils"
)

const TEST_ROW int = 1
const TEST_COLUMN int = 1

func AssertValue(t *testing.T, actual int, expected int) {
	if expected != actual {
		t.Errorf("unexpected value. expected: %d, actual: %d", expected, actual)
	}
}

func assertPencilMarks(t *testing.T, actual []int, expected []int) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("unexpected pencil_marks. expected: %v, actual: %v", expected, actual)
	}
}

func TestGetValueDefault(t *testing.T) {
	cell, err := NewCell(TEST_ROW, TEST_COLUMN)

	AssertNoError(t, err)
	AssertValue(t, cell.GetValue(), Empty)
}

func TestNewCellInvalidRowOrColumn(t *testing.T) {
	invalid_row_and_column_ids := [][2]int{
		{0, 1},
		{1, 0},
		{10, 1},
		{1, 10},
	}

	for _, invalid_row_and_column := range invalid_row_and_column_ids {
		row := invalid_row_and_column[0]
		column := invalid_row_and_column[1]

		t.Run(fmt.Sprintf("%d,%d", row, column), func(t *testing.T) {
			cell, err := NewCell(row, column)

			AssertError(t, err)
			if cell != nil {
				t.Errorf("unexpected cell")
			}
		})
	}
}

func TestGetRowIdAndGetColumnId(t *testing.T) {
	for row := 1; row <= 9; row++ {
		for column := 1; column <= 9; column++ {
			t.Run(fmt.Sprintf("%d,%d", row, column), func(t *testing.T) {
				cell, err := NewCell(row, column)

				AssertNoError(t, err)
				if cell == nil {
					t.Errorf("missing cell")
				}

				actual_row_id := cell.GetRowId()
				if actual_row_id != row {
					t.Errorf("unexpected row ID. expected: %d, actual: %d", row, actual_row_id)
				}

				actual_column_id := cell.GetColumnId()
				if actual_column_id != column {
					t.Errorf("unexpected column ID. expected: %d, actual: %d", column, actual_column_id)
				}

			})
		}
	}
}

func TestGetBoxId(t *testing.T) {
	expected_box_ids := []int{
		1, 1, 1, 2, 2, 2, 3, 3, 3,
		1, 1, 1, 2, 2, 2, 3, 3, 3,
		1, 1, 1, 2, 2, 2, 3, 3, 3,
		4, 4, 4, 5, 5, 5, 6, 6, 6,
		4, 4, 4, 5, 5, 5, 6, 6, 6,
		4, 4, 4, 5, 5, 5, 6, 6, 6,
		7, 7, 7, 8, 8, 8, 9, 9, 9,
		7, 7, 7, 8, 8, 8, 9, 9, 9,
		7, 7, 7, 8, 8, 8, 9, 9, 9,
	}

	i := 0
	for row := 1; row <= 9; row++ {
		for column := 1; column <= 9; column++ {
			t.Run(fmt.Sprintf("%d,%d", row, column), func(t *testing.T) {
				cell, err := NewCell(row, column)

				AssertNoError(t, err)
				if cell == nil {
					t.Errorf("missing cell")
				}

				expected_box_id := expected_box_ids[i]

				actual_box_id := cell.GetBoxId()
				if actual_box_id != expected_box_id {
					t.Errorf("unexpected Nox ID. expected: %d, actual: %d", expected_box_id, actual_box_id)
				}
			})

			i++
		}
	}
}

func TestSetValue(t *testing.T) {
	for value_to_set := 1; value_to_set <= 9; value_to_set++ {
		t.Run(strconv.Itoa(value_to_set), func(t *testing.T) {
			cell, _ := NewCell(TEST_ROW, TEST_COLUMN)

			AssertNoError(t, cell.SetValue(value_to_set))
			AssertValue(t, cell.GetValue(), value_to_set)
		})
	}
}

func TestSetValueEmpty(t *testing.T) {
	cell, err := NewCell(TEST_ROW, TEST_COLUMN)
	AssertNoError(t, err)

	AssertNoError(t, cell.SetValue(1))
	AssertNoError(t, cell.SetValue(Empty))
	AssertValue(t, cell.GetValue(), Empty)
}

func TestSetValueIdempotency(t *testing.T) {
	cell, err := NewCell(TEST_ROW, TEST_COLUMN)
	AssertNoError(t, err)

	AssertNoError(t, cell.SetValue(1))
	AssertNoError(t, cell.SetValue(1))
	AssertValue(t, cell.GetValue(), 1)
}

func TestSetValueInvalid(t *testing.T) {
	values_to_set := []int{0, 10}
	for _, value_to_set := range values_to_set {
		t.Run(strconv.Itoa(value_to_set), func(t *testing.T) {
			cell, _ := NewCell(TEST_ROW, TEST_COLUMN)

			AssertError(t, cell.SetValue(value_to_set))
			AssertValue(t, cell.GetValue(), Empty)
		})
	}
}

func TestGetPencilMarksDefault(t *testing.T) {
	cell, _ := NewCell(TEST_ROW, TEST_COLUMN)

	assertPencilMarks(t, cell.GetPencilMarks(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func emptyPencilMarks(cell *Cell) {
	for pencil_mark := 1; pencil_mark <= 9; pencil_mark++ {
		cell.RemovePencilMark(pencil_mark)
	}
}

func TestAddPencilMark(t *testing.T) {
	cell, _ := NewCell(TEST_ROW, TEST_COLUMN)
	emptyPencilMarks(cell)

	expected_pencil_marks := []int{}

	for pencil_mark_to_add := 1; pencil_mark_to_add <= 9; pencil_mark_to_add++ {
		expected_pencil_marks = append(expected_pencil_marks, pencil_mark_to_add)

		AssertNoError(t, cell.AddPencilMark(pencil_mark_to_add))
		assertPencilMarks(t, cell.GetPencilMarks(), expected_pencil_marks)
	}
}

func TestAddPencilMarkIdempotency(t *testing.T) {
	cell, _ := NewCell(TEST_ROW, TEST_COLUMN)
	emptyPencilMarks(cell)

	AssertNoError(t, cell.AddPencilMark(1))
	AssertNoError(t, cell.AddPencilMark(1))
	assertPencilMarks(t, cell.GetPencilMarks(), []int{1})
}

func TestAddPencilMarkInvalid(t *testing.T) {
	pencil_marks_to_add := []int{0, 10}

	for _, pencil_mark_to_add := range pencil_marks_to_add {
		t.Run(strconv.Itoa(pencil_mark_to_add), func(t *testing.T) {
			cell, _ := NewCell(TEST_ROW, TEST_COLUMN)
			emptyPencilMarks(cell)

			AssertError(t, cell.AddPencilMark(pencil_mark_to_add))
			assertPencilMarks(t, cell.GetPencilMarks(), []int{})
		})
	}
}

func TestRemovePencilMark(t *testing.T) {
	cell, _ := NewCell(TEST_ROW, TEST_COLUMN)

	expected_pencil_marks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for pencil_mark_to_remove := 1; pencil_mark_to_remove <= 9; pencil_mark_to_remove++ {
		expected_pencil_marks = expected_pencil_marks[1:]

		AssertNoError(t, cell.RemovePencilMark(pencil_mark_to_remove))
		assertPencilMarks(t, cell.GetPencilMarks(), expected_pencil_marks)
	}
}

func TestRemovePencilMarkIdempotency(t *testing.T) {
	cell, _ := NewCell(TEST_ROW, TEST_COLUMN)

	expected_pencil_marks := []int{2, 3, 4, 5, 6, 7, 8, 9}

	AssertNoError(t, cell.RemovePencilMark(1))
	AssertNoError(t, cell.RemovePencilMark(1))
	assertPencilMarks(t, cell.GetPencilMarks(), expected_pencil_marks)
}

func TestRemovePencilMarkInvalid(t *testing.T) {
	pencil_marks_to_remove := []int{0, 10}

	for _, pencil_mark_to_remove := range pencil_marks_to_remove {
		t.Run(strconv.Itoa(pencil_mark_to_remove), func(t *testing.T) {
			cell, _ := NewCell(TEST_ROW, TEST_COLUMN)

			expected_pencil_marks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

			AssertError(t, cell.RemovePencilMark(pencil_mark_to_remove))
			assertPencilMarks(t, cell.GetPencilMarks(), expected_pencil_marks)
		})
	}
}

func TestAddPencilMarks(t *testing.T) {
	cell, _ := NewCell(TEST_ROW, TEST_COLUMN)
	emptyPencilMarks(cell)

	pencil_marks_to_add := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	AssertNoError(t, cell.AddPencilMarks(pencil_marks_to_add))
	assertPencilMarks(t, cell.GetPencilMarks(), pencil_marks_to_add)
}

func TestAddPencilMarksIdempotency(t *testing.T) {
	cell, _ := NewCell(TEST_ROW, TEST_COLUMN)
	emptyPencilMarks(cell)

	pencil_marks_to_add := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	AssertNoError(t, cell.AddPencilMarks(pencil_marks_to_add))
	AssertNoError(t, cell.AddPencilMarks(pencil_marks_to_add))
	assertPencilMarks(t, cell.GetPencilMarks(), pencil_marks_to_add)
}

func TestAddPencilMarksInvalid(t *testing.T) {
	pencil_mark_arrays_to_add := [][]int{
		{1, 0},
		{1, 10},
	}

	for _, pencil_marks_to_add := range pencil_mark_arrays_to_add {
		t.Run(strconv.Itoa(pencil_marks_to_add[1]), func(t *testing.T) {
			cell, _ := NewCell(TEST_ROW, TEST_COLUMN)
			emptyPencilMarks(cell)

			AssertError(t, cell.AddPencilMarks(pencil_marks_to_add))
			assertPencilMarks(t, cell.GetPencilMarks(), []int{})
		})
	}
}

func TestRemovePencilMarks(t *testing.T) {
	cell, _ := NewCell(TEST_ROW, TEST_COLUMN)

	pencil_marks_to_remove := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	AssertNoError(t, cell.RemovePencilMarks(pencil_marks_to_remove))
	assertPencilMarks(t, cell.GetPencilMarks(), []int{})
}

func TestRemovePencilMarksIdempotency(t *testing.T) {
	cell, _ := NewCell(TEST_ROW, TEST_COLUMN)

	pencil_marks_to_remove := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	AssertNoError(t, cell.RemovePencilMarks(pencil_marks_to_remove))
	AssertNoError(t, cell.RemovePencilMarks(pencil_marks_to_remove))
	assertPencilMarks(t, cell.GetPencilMarks(), []int{})
}

func TestRemovePencilMarksInvalid(t *testing.T) {
	pencil_mark_arrays_to_remove := [][]int{
		{1, 0},
		{1, 10},
	}

	for _, pencil_marks_to_remove := range pencil_mark_arrays_to_remove {
		t.Run(strconv.Itoa(pencil_marks_to_remove[1]), func(t *testing.T) {
			cell, _ := NewCell(TEST_ROW, TEST_COLUMN)

			expected_pencil_marks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

			AssertError(t, cell.RemovePencilMarks(pencil_marks_to_remove))
			assertPencilMarks(t, cell.GetPencilMarks(), expected_pencil_marks)
		})
	}
}
