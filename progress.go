package ewwgo

import (
	"fmt"
)

type Progress struct {
	Flipped     bool
	Orientation string
	Value       float64

	*General
}

func (m *Progress) String() string {
	var attr []string
	attr = append(attr, fmt.Sprintf(":orientation '%s' ", m.Orientation))
	attr = append(attr, fmt.Sprintf(":flipped %t ", m.Flipped))
	attr = append(attr, fmt.Sprintf(":value %f ", m.Value))
	attr = append(attr, m.General.String()...)

	return fmt.Sprintf("(progress %s)", stringBuilder(attr))
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

func (m *Progress) SetFlipped(t bool) *Progress {
	m.Flipped = t

	return m
}

func (m *Progress) SetOrientation(s string) *Progress {
	if s != "v" && s != "h" && s != "vertical" && s != "horizontal" {
		return m
	}

	m.Orientation = s

	return m
}

func (m *Progress) SetValue(f float64) *Progress {
	if f > 100 {
		return m
	}
	m.Value = f

	return m
}

func (m *Progress) SetGeneral(g *General) *Progress {
	m.General = g

	return m
}
