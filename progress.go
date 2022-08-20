package ewwgo

import (
	"errors"
	"fmt"
)

type Progress struct {
	Flipped     bool
	Orientation string
	Value       float64

	*General
}

func (p Progress) String() string {
	class := fmt.Sprintf(":class '%s' ", p.Class)
	flipped := fmt.Sprintf(":flipped %t ", p.Flipped)
	orientation := fmt.Sprintf(":orientation '%s' ", p.Orientation)
	value := fmt.Sprintf(":value %f ", p.Value)
	active := fmt.Sprintf(":active %t ", p.Active)

	return fmt.Sprintf("(progress %s %s %s %s %s)", active, class, flipped, orientation, value)

}

func NewProgress() *Progress {
	g := NewGeneral()

	p := &Progress{
		Flipped:     false,
		Orientation: "h",
		Value:       0,

		General: g,
	}

	return p
}

func (m *Progress) SetFlipped(t bool) error {
	m.Flipped = t

	return nil
}

func (m *Progress) SetOrientation(s string) error {
	if s != "v" && s != "h" && s != "vertical" && s != "horizontal" {
		return errors.New("err: should be v or h or vertical or horizontal")
	}

	m.Orientation = s

	return nil
}

func (m *Progress) SetValue(f float64) error {
	m.Value = f

	return nil
}

func (m *Progress) SetGeneral(g *General) error {
	m.General = g

	return nil
}
