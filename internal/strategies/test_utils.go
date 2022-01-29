package strategies

import (
	"testing"

	"github.com/alltilla/sudoku-solver/internal/sudoku"
	"github.com/alltilla/sudoku-solver/internal/test_utils"
)

func AssertGridEqualsWithPrettyString(t *testing.T, actual_grid *sudoku.Grid, expected_pretty_string string) {
	expected_grid := sudoku.NewGrid()
	test_utils.AssertNoError(t, expected_grid.LoadPrettyString(expected_pretty_string))

	if !actual_grid.Equals(expected_grid) {
		t.Errorf("unexpected grid")
	}
}

func AssertChanged(t *testing.T, changed bool) {
	if !changed {
		t.Errorf("missing change")
	}
}

func AssertNoChanged(t *testing.T, changed bool) {
	if changed {
		t.Errorf("unexpected change")
	}
}
