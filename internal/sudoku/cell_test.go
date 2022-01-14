package sudoku

import (
	"reflect"
	"strconv"
	"testing"

	. "github.com/alltilla/sudoku-solver/internal/test_utils"
)

func assertValue(t *testing.T, actual int, expected int) {
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
	cell := NewCell()

	assertValue(t, cell.GetValue(), Empty)
}

func TestSetValue(t *testing.T) {
	for value_to_set := 1; value_to_set <= 9; value_to_set++ {
		t.Run(strconv.Itoa(value_to_set), func(t *testing.T) {
			cell := NewCell()

			AssertNoError(t, cell.SetValue(value_to_set))
			assertValue(t, cell.GetValue(), value_to_set)
		})
	}
}

func TestSetValueEmpty(t *testing.T) {
	cell := NewCell()

	AssertNoError(t, cell.SetValue(1))
	AssertNoError(t, cell.SetValue(Empty))
	assertValue(t, cell.GetValue(), Empty)
}

func TestSetValueIdempotency(t *testing.T) {
	cell := NewCell()

	AssertNoError(t, cell.SetValue(1))
	AssertNoError(t, cell.SetValue(1))
	assertValue(t, cell.GetValue(), 1)
}

func TestSetValueInvalid(t *testing.T) {
	values_to_set := []int{0, 10}
	for _, value_to_set := range values_to_set {
		t.Run(strconv.Itoa(value_to_set), func(t *testing.T) {
			cell := NewCell()

			AssertError(t, cell.SetValue(value_to_set))
			assertValue(t, cell.GetValue(), Empty)
		})
	}
}

func TestGetPencilMarksDefault(t *testing.T) {
	cell := NewCell()

	assertPencilMarks(t, cell.GetPencilMarks(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func emptyPencilMarks(cell *Cell) {
	for pencil_mark := 1; pencil_mark <= 9; pencil_mark++ {
		cell.RemovePencilMark(pencil_mark)
	}
}

func TestAddPencilMark(t *testing.T) {
	cell := NewCell()
	emptyPencilMarks(cell)

	expected_pencil_marks := []int{}

	for pencil_mark_to_add := 1; pencil_mark_to_add <= 9; pencil_mark_to_add++ {
		expected_pencil_marks = append(expected_pencil_marks, pencil_mark_to_add)

		AssertNoError(t, cell.AddPencilMark(pencil_mark_to_add))
		assertPencilMarks(t, cell.GetPencilMarks(), expected_pencil_marks)
	}
}

func TestAddPencilMarkIdempotency(t *testing.T) {
	cell := NewCell()
	emptyPencilMarks(cell)

	AssertNoError(t, cell.AddPencilMark(1))
	AssertNoError(t, cell.AddPencilMark(1))
	assertPencilMarks(t, cell.GetPencilMarks(), []int{1})
}

func TestAddPencilMarkInvalid(t *testing.T) {
	pencil_marks_to_add := []int{0, 10}

	for _, pencil_mark_to_add := range pencil_marks_to_add {
		t.Run(strconv.Itoa(pencil_mark_to_add), func(t *testing.T) {
			cell := NewCell()
			emptyPencilMarks(cell)

			AssertError(t, cell.AddPencilMark(pencil_mark_to_add))
			assertPencilMarks(t, cell.GetPencilMarks(), []int{})
		})
	}
}

func TestRemovePencilMark(t *testing.T) {
	cell := NewCell()

	expected_pencil_marks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for pencil_mark_to_remove := 1; pencil_mark_to_remove <= 9; pencil_mark_to_remove++ {
		expected_pencil_marks = expected_pencil_marks[1:]

		AssertNoError(t, cell.RemovePencilMark(pencil_mark_to_remove))
		assertPencilMarks(t, cell.GetPencilMarks(), expected_pencil_marks)
	}
}

func TestRemovePencilMarkIdempotency(t *testing.T) {
	cell := NewCell()

	expected_pencil_marks := []int{2, 3, 4, 5, 6, 7, 8, 9}

	AssertNoError(t, cell.RemovePencilMark(1))
	AssertNoError(t, cell.RemovePencilMark(1))
	assertPencilMarks(t, cell.GetPencilMarks(), expected_pencil_marks)
}

func TestRemovePencilMarkInvalid(t *testing.T) {
	pencil_marks_to_remove := []int{0, 10}

	for _, pencil_mark_to_remove := range pencil_marks_to_remove {
		t.Run(strconv.Itoa(pencil_mark_to_remove), func(t *testing.T) {
			cell := NewCell()

			expected_pencil_marks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

			AssertError(t, cell.RemovePencilMark(pencil_mark_to_remove))
			assertPencilMarks(t, cell.GetPencilMarks(), expected_pencil_marks)
		})
	}
}
