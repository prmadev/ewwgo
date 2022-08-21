package ewwgo

import (
	"errors"
	"fmt"
)

type Box struct {
	Spacing     int
	Orientation string
	SpaceEvenly bool
	Content     []string

	*General
}

func (m *Box) String() string {

	var c string
	for _, v := range m.Content {
		c = c + v
	}
	return fmt.Sprintf("(box :class '%s' :spacing %d :orientation '%s' :space-evenly %t %s)", m.Class, m.Spacing, m.Orientation, m.SpaceEvenly, c)
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

func (m *Box) SetSpacing(d int) error {
	m.Spacing = d
	return nil
}

func (m *Box) SetOrientation(s string) error {
	if s != "v" && s != "h" && s != "vertical" && s != "horizontal" {
		return errors.New("err: should be v or h or vertical or horizontal")
	}

	m.Orientation = s

	return nil
}

func (m *Box) SetSpaceEvenly(t bool) error {
	m.SpaceEvenly = t

	return nil
}


func (m *Box) AppendContent(w string ) error {
	m.Content = append(m.Content, w) 


	return nil
}

func (m *Box) SetGeneral(g *General) error {
	m.General = g

	return nil
}
