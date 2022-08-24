package ewwgo

import (
	"fmt"
)

type Box struct {
	Spacing     int
	Orientation string
	SpaceEvenly bool
	Content     []Widget

	*General
}

func (m *Box) String() string {
	var attr []string
	attr = append(attr, fmt.Sprintf(":spacing %d", m.Spacing))
	attr = append(attr, fmt.Sprintf(":orientation '%s'", m.Orientation))
	attr = append(attr, fmt.Sprintf(":space-evenly '%t'", m.SpaceEvenly))
	attr = append(attr, m.General.String()...)
	for _, w := range m.Content {
		attr = append(attr, w.String())
	}

	return fmt.Sprintf("(box %s)", stringBuilder(attr))
}

func NewBox() *Box {
	g := NewGeneral()

	b := &Box{
		Spacing:     0,
		Orientation: "h",
		SpaceEvenly: false,

		General: g,
	}

	return b
}

func (m *Box) SetSpacing(d int) *Box {
	m.Spacing = d
	return m
}

func (m *Box) SetOrientation(s string) *Box {
	if s != "v" && s != "h" && s != "vertical" && s != "horizontal" {
		return m
	}

	m.Orientation = s

	return m
}

func (m *Box) SetSpaceEvenly(t bool) *Box {
	m.SpaceEvenly = t

	return m
}

func (m *Box) AppendContent(ws ...Widget) *Box {
	m.Content = append(m.Content, ws...)

	return m
}

func (m *Box) SetGeneral(g *General) *Box {
	m.General = g

	return m
}
