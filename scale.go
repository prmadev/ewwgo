package ewwgo

import (
	"errors"
	"fmt"
)

type Scale struct {
	Flipped     bool
	Orientation string
	Value       float64
	DrawValue   bool
	Marks       string
	Min         float64
	Max         float64
	Timeout     string
	Onchange    string

	*General
}

func (m Scale) String() string {
	var attr []string
	attr = append(attr, fmt.Sprintf(":flipped %t ", m.Flipped))
	attr = append(attr, fmt.Sprintf(":orientation '%s' ", m.Orientation))
	attr = append(attr, fmt.Sprintf(":value %f ", m.Value))
	// attr = append(attr, fmt.Sprintf(":marks '%s' ", s.Marks)) //unstable
	attr = append(attr, fmt.Sprintf(":draw-value %t ", m.DrawValue))
	attr = append(attr, fmt.Sprintf(":min %f ", m.Min))
	attr = append(attr, fmt.Sprintf(":max %f ", m.Max))
	// attr = append(attr, fmt.Sprintf(":timeout %s ", s.Timeout))
	attr = append(attr, fmt.Sprintf(":onchange '%s' ", m.Onchange))
	attr = append(attr, m.General.String()...)

	str := stringBuilder(attr)

	return fmt.Sprintf("(scale %s)", str)
}

func NewScale() *Scale {
	g := NewGeneral()

	s := &Scale{
		Flipped:     false,
		Orientation: "h",
		Value:       0,
		DrawValue:   false,
		Marks:       "",
		Min:         0,
		Max:         100,
		Timeout:     "10s",
		Onchange:    "",

		General: g,
	}

	return s
}

func (m *Scale) SetGeneral(g *General) error {
	m.General = g

	return nil
}

func (m *Scale) SetFlipped(t bool) error {
	m.Flipped = t

	return nil
}

func (m *Scale) SetDrawValue(t bool) error {
	m.DrawValue = t

	return nil
}

func (m *Scale) SetValue(f float64) error {
	m.Value = f

	return nil
}

func (m *Scale) SetOrientation(s string) error {
	if s != "v" && s != "h" && s != "vertical" && s != "horizontal" {
		return errors.New("err: should be v or h or vertical or horizontal")
	}

	m.Orientation = s

	return nil
}

func (m *Scale) SetMarks(s string) error {
	m.Marks = s

	return nil
}

func (m *Scale) SetMin(f float64) error {
	m.Min = f

	return nil
}

func (m *Scale) SetMax(f float64) error {
	m.Max = f

	return nil
}

func (m *Scale) SetTimeout(s string) error {
	m.Timeout = s

	return nil
}


func (m *Scale) SetOnChange(s string) error {
	m.Onchange = s

	return nil
}

