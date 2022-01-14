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
