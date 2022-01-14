package sudoku

import (
	"testing"
)

func assertValue(t *testing.T, actual int, expected int) {
	if expected != actual {
		t.Errorf("unexpected value. expected: %d, actual: %d", expected, actual)
	}
}

func TestGetValueDefault(t *testing.T) {
	cell := NewCell()

	assertValue(t, cell.GetValue(), Empty)
}
