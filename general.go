package ewwgo

import (
	"errors"
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
	ws = append(ws, fmt.Sprintf(":width %d ", m.Width))
	ws = append(ws, fmt.Sprintf(":height %d ", m.Height))
	ws = append(ws, fmt.Sprintf(":active %t ", m.Active))
	ws = append(ws, fmt.Sprintf(":tooltip '%s '", m.Tooltip))
	ws = append(ws, fmt.Sprintf(":visible %t ", m.Visible))
	ws = append(ws, fmt.Sprintf(":style '%s' ", m.Style))

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

func (m *General) SetClass(s string) error {
	m.Class = s

	return nil
}

func (m *General) SetValign(s string) error {
	if s != "fill" && s != "baseline" && s != "center" && s != "start" && s != "end" {
		return errors.New("err: value should be either fill, baseline, center, start or end")
	}

	m.Valign = s

	return nil
}

func (m *General) SetXalign(s string) error {
	if s != "fill" && s != "baseline" && s != "center" && s != "start" && s != "end" {
		return errors.New("err: value should be either fill, baseline, center, start or end")
	}

	m.Halign = s

	return nil
}

func (m *General) SetHexpand(t bool) error {
	m.Hexpand = t

	return nil
}

func (m *General) SetVexpand(t bool) error {
	m.Vexpand = t

	return nil
}

func (m *General) SetWidth(d int) error {
	m.Width = d

	return nil
}

func (m *General) SetHeight(d int) error {
	m.Height = d

	return nil
}

func (m *General) SetActive(t bool) error {
	m.Active = t

	return nil
}

func (m *General) SetTooltip(s string) error {
	m.Tooltip = s

	return nil
}

func (m *General) SetVisible(b bool) error {
	m.Visible = b

	return nil
}

func (m *General) SetStyle(s map[string]string) error {
	m.Style = s

	return nil
}

func (m *General) ModifyStyle(k, v string) error {
	m.Style[k] = v

	return nil
}
