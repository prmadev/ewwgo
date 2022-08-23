package ewwgo

import (
	"fmt"
)

type General struct {
	Class   string
	Valign  string
	Halign  string
	Vexpand bool
	Hexpand bool
	Width   int
	Height  int
	Active  bool
	Tooltip string
	Visible bool
	Style   Css
}

func (m General) String() []string {
	var ws []string

	ws = append(ws, fmt.Sprintf(":class '%s' ", m.Class))

	ws = append(ws, fmt.Sprintf(":valign '%s' ", m.Valign))
	ws = append(ws, fmt.Sprintf(":halign '%s' ", m.Halign))

	ws = append(ws, fmt.Sprintf(":vexpand %t ", m.Vexpand))
	ws = append(ws, fmt.Sprintf(":hexpand %t ", m.Hexpand))

	ws = append(ws, fmt.Sprintf(":width %d ", m.Width))
	ws = append(ws, fmt.Sprintf(":height %d ", m.Height))

	ws = append(ws, fmt.Sprintf(":active %t ", m.Active))
	ws = append(ws, fmt.Sprintf(":visible %t ", m.Visible))

	ws = append(ws, fmt.Sprintf(":tooltip '%s '", m.Tooltip))
	ws = append(ws, fmt.Sprintf(":style '%s' ", m.Style.String()))

	return ws
}

func NewGeneral() *General {
	m := make(Css)

	g := &General{
		Class:   "",
		Valign:  "center",
		Halign:  "center",
		Vexpand: false,
		Hexpand: false,
		Width:   0,
		Height:  0,
		Active:  false,
		Tooltip: "",
		Visible: true,
		Style:   m,
	}

	return g
}

func (m *General) SetClass(s string) *General {
	m.Class = s

	return m
}

func (m *General) SetValign(s string) *General {
	if s != "fill" && s != "baseline" && s != "center" && s != "start" && s != "end" {
		return m
	}

	m.Valign = s

	return m
}

func (m *General) SetHalign(s string) *General {
	if s != "fill" && s != "baseline" && s != "center" && s != "start" && s != "end" {
		return m
	}

	m.Halign = s

	return m
}

func (m *General) SetHexpand(t bool) *General {
	m.Hexpand = t

	return m
}

func (m *General) SetVexpand(t bool) *General {
	m.Vexpand = t

	return m
}

func (m *General) SetWidth(d int) *General {
	m.Width = d

	return m
}

func (m *General) SetHeight(d int) *General {
	m.Height = d

	return m
}

func (m *General) SetActive(t bool) *General {
	m.Active = t

	return m
}

func (m *General) SetTooltip(s string) *General {
	m.Tooltip = s

	return m
}

func (m *General) SetVisible(b bool) *General {
	m.Visible = b

	return m
}

func (m *General) SetStyle(s map[string]string) *General {
	m.Style = s

	return m
}

func (m *General) ModifyStyle(k, v string) *General {
	m.Style[k] = v

	return m
}
