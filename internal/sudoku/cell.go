package sudoku

import "fmt"

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

func checkDigitValidity(digit int) error {
	if digit > 9 {
		return fmt.Errorf("digit cannot be larger than 9: %d", digit)
	}
	if digit < 1 {
		return fmt.Errorf("digit cannot be smaller than 1: %d", digit)
	}
	return nil
}

func (c *Cell) GetValue() int {
	return c.value
}

func (c *Cell) SetValue(value int) error {
	if value != Empty {
		err := checkDigitValidity(value)
		if err != nil {
			return err
		}
	}

	c.value = value
	return nil
}
