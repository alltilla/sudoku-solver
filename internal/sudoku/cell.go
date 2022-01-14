package sudoku

import "fmt"

const Empty int = -127

type Cell struct {
	value        int
	pencil_marks [9]bool
}

func NewCell() *Cell {
	value := Empty
	pencil_marks := [9]bool{true, true, true, true, true, true, true, true, true}

	c := Cell{
		value,
		pencil_marks,
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

func (c *Cell) GetPencilMarks() []int {
	pencil_marks := []int{}
	for index, set := range c.pencil_marks {
		if set {
			digit := index + 1
			pencil_marks = append(pencil_marks, digit)
		}
	}

	return pencil_marks
}
