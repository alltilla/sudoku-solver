package sudoku

import (
	"fmt"
	"reflect"
)

const Empty int = -127

type Cell struct {
	row          int
	column       int
	value        int
	pencil_marks [9]bool
}

func CheckRowAndColumnValidity(row int, column int) error {
	if row > 9 {
		return fmt.Errorf("row is larger than 9: %d", row)
	}
	if row < 1 {
		return fmt.Errorf("row is smaller than 1: %d", row)
	}
	if column > 9 {
		return fmt.Errorf("column is larger than 9: %d", column)
	}
	if column < 1 {
		return fmt.Errorf("column is smaller than 1: %d", column)
	}

	return nil
}

func CheckBoxValidity(box int) error {
	if box > 9 {
		return fmt.Errorf("box is larger than 9: %d", box)
	}
	if box < 1 {
		return fmt.Errorf("box is smaller than 1: %d", box)
	}
	return nil
}

func NewCell(row int, column int) (*Cell, error) {
	err := CheckRowAndColumnValidity(row, column)
	if err != nil {
		return nil, err
	}

	value := Empty
	pencil_marks := [9]bool{true, true, true, true, true, true, true, true, true}

	c := Cell{
		row,
		column,
		value,
		pencil_marks,
	}

	return &c, nil
}

func (c *Cell) Equals(other *Cell) bool {
	return c.row == other.row &&
		c.column == other.column &&
		c.value == other.value &&
		reflect.DeepEqual(c.pencil_marks, other.pencil_marks)
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

func (c *Cell) GetRowId() int {
	return c.row
}

func (c *Cell) GetColumnId() int {
	return c.column
}

func (c *Cell) GetValue() int {
	return c.value
}

func (c *Cell) GetBoxId() int {
	if c.GetRowId() <= 3 {
		if c.GetColumnId() <= 3 {
			return 1
		} else if c.GetColumnId() <= 6 {
			return 2
		} else {
			return 3
		}
	} else if c.GetRowId() <= 6 {
		if c.GetColumnId() <= 3 {
			return 4
		} else if c.GetColumnId() <= 6 {
			return 5
		} else {
			return 6
		}
	} else {
		if c.GetColumnId() <= 3 {
			return 7
		} else if c.GetColumnId() <= 6 {
			return 8
		} else {
			return 9
		}
	}
}

func (c *Cell) SetValue(value int) error {
	if value == Empty {
		c.AddPencilMarks([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		c.value = value

		return nil
	}

	err := checkDigitValidity(value)
	if err != nil {
		return err
	}

	c.value = value
	c.RemovePencilMarks([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

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

func (c *Cell) changePencilMark(pencil_mark int, set bool) error {
	err := checkDigitValidity(pencil_mark)
	if err != nil {
		return err
	}

	c.pencil_marks[pencil_mark-1] = set
	return nil
}

func (c *Cell) AddPencilMark(pencil_mark int) error {
	return c.changePencilMark(pencil_mark, true)
}

func (c *Cell) RemovePencilMark(pencil_mark int) error {
	return c.changePencilMark(pencil_mark, false)
}

func (c *Cell) changePencilMarks(pencil_marks []int, set bool) error {
	for _, pencil_mark := range pencil_marks {
		err := checkDigitValidity(pencil_mark)
		if err != nil {
			return err
		}
	}

	for _, pencil_mark := range pencil_marks {
		c.pencil_marks[pencil_mark-1] = set
	}

	return nil
}

func (c *Cell) AddPencilMarks(pencil_marks []int) error {
	return c.changePencilMarks(pencil_marks, true)
}

func (c *Cell) RemovePencilMarks(pencil_marks []int) error {
	return c.changePencilMarks(pencil_marks, false)
}
