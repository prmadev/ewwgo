package ewwgo

import "fmt"

type Expander struct {
	Name     string
	Expanded bool
	Content  []Widget

	*General
}

func (m *Expander) String() string {
	var attr []string
	attr = append(attr, fmt.Sprintf(":name '%s'", m.Name))
	attr = append(attr, fmt.Sprintf(":expanded %t", m.Expanded))
	attr = append(attr, m.General.String()...)
	for _, w := range m.Content {
		attr = append(attr, w.String())
	}

	return fmt.Sprintf("(expander %s)", stringBuilder(attr))
}

func NewExpander() *Expander {
	g := NewGeneral()

	b := &Expander{
		Name:     "",
		Expanded: false,

		General: g,
	}

	return b
}

func (m *Expander) SetExpanded(t bool) *Expander {
	m.Expanded = t

	return m
}

func (m *Expander) SetName(name string) *Expander {
	m.Name = name

	return m
}

func (m *Expander) AppendContent(ws ...Widget) *Expander {
	m.Content = append(m.Content, ws...)

	return m
}

func (m *Expander) SetGeneral(g *General) *Expander {
	m.General = g

	return m
}
