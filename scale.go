package ewwgo

import (
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
	RoundDigits int

	*General
}

func (m Scale) String() string {
	var attr []string
	attr = append(attr, fmt.Sprintf(":flipped %t ", m.Flipped))
	attr = append(attr, fmt.Sprintf(":orientation '%s' ", m.Orientation))
	attr = append(attr, fmt.Sprintf(":value %f ", m.Value))
	attr = append(attr, fmt.Sprintf(":marks '%s' ", m.Marks)) // unstable
	attr = append(attr, fmt.Sprintf(":draw-value %t ", m.DrawValue))
	attr = append(attr, fmt.Sprintf(":min %f ", m.Min))
	attr = append(attr, fmt.Sprintf(":max %f ", m.Max))
	attr = append(attr, fmt.Sprintf(":timeout '%s' ", m.Timeout))
	attr = append(attr, fmt.Sprintf(":onchange '%s' ", m.Onchange))
	attr = append(attr, fmt.Sprintf(":round-digits %d ", m.RoundDigits))
	attr = append(attr, m.General.String()...)

	return fmt.Sprintf("(scale %s)", stringBuilder(attr))
}

func NewScale() *Scale {
	g := NewGeneral()

	s := &Scale{
		Flipped:     false,
		Orientation: "h",
		Value:       0,
		DrawValue:   false,
		Marks:       " ",
		Min:         0,
		Max:         100,
		Timeout:     "10s",
		Onchange:    " ",
		RoundDigits: 1,

		General: g,
	}

	return s
}

func (m *Scale) SetGeneral(g *General) *Scale {
	m.General = g

	return nil
}

func (m *Scale) SetFlipped(t bool) *Scale {
	m.Flipped = t

	return nil
}

func (m *Scale) SetDrawValue(t bool) *Scale {
	m.DrawValue = t

	return nil
}

func (m *Scale) SetValue(f float64) *Scale {
	m.Value = f

	return m
}

func (m *Scale) SetRoundDigits(d int) *Scale {
	m.RoundDigits = d

	return m
}
func (m *Scale) SetOrientation(s string) *Scale {
	if s != "v" && s != "h" && s != "vertical" && s != "horizontal" {
		return m
	}

	m.Orientation = s

	return m
}

func (m *Scale) SetMarks(s string) *Scale {
	m.Marks = s

	return m
}

func (m *Scale) SetMin(f float64) *Scale {
	m.Min = f

	return m
}

func (m *Scale) SetMax(f float64) *Scale {
	m.Max = f

	return m
}

func (m *Scale) SetTimeout(s string) *Scale {
	m.Timeout = s

	return m
}

func (m *Scale) SetOnChange(s string) *Scale {
	m.Onchange = s

	return m
}
