package sudoku

type Grid struct {
	cells [9][9]*Cell
}

func NewGrid() *Grid {
	g := Grid{}

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			g.cells[row][column] = NewCell()
		}
	}

	return &g
}
