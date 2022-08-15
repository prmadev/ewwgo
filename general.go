package ewwgo

import "fmt"

type Widget struct {
	Class   string
	Valign  string
	Halign  string
	Vexpand bool
	Hexpand bool
	Width   int
	Heigth  int
	Active  bool
	Tooltip string
	Visible bool
	Style   string
}

func (w Widget) Printer() []string {
	var ws []string

	ws = append(ws, fmt.Sprintf(":class '%s'", w.Class))
	ws = append(ws, fmt.Sprintf(":valign '%s'", w.Valign))
	ws = append(ws, fmt.Sprintf(":halign '%s'", w.Halign))
	ws = append(ws, fmt.Sprintf(":width %d", w.Width))
	ws = append(ws, fmt.Sprintf(":heigth %d", w.Heigth))
	ws = append(ws, fmt.Sprintf(":active %t", w.Active))
	ws = append(ws, fmt.Sprintf(":tooltip '%s'", w.Tooltip))
	ws = append(ws, fmt.Sprintf(":visible %t", w.Visible))
	ws = append(ws, fmt.Sprintf(":style '%s'", w.Style))

	return ws
}
