package sudoku

const Empty int = -127

type Cell struct {
	value int
}

func NewCell() *Cell {
	value := Empty

	c := Cell{
		value,
	}

	return &c
}

func (c *Cell) GetValue() int {
	return c.value
}
