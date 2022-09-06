package ewwgo

import "fmt"

type (
	OrientationType string
	AlignType       string
)

const (
	VERTICAL   OrientationType = "vertical"
	HORIZONTAL OrientationType = "horizontal"
	FILL       AlignType       = "fill"
	BASELINE   AlignType       = "baseline"
	CENTER     AlignType       = "center"
	START      AlignType       = "start"
	END        AlignType       = "end"
)

func Spacing(in int) OptFunc {
	return func(wid *Widget) {
		if wid.Type != BOX {
			return
		}
		key := "spacing"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func SpaceEvenly(in bool) OptFunc {
	return func(wid *Widget) {
		if wid.Type != BOX {
			return
		}
		key := "space-evenly"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Class(in string) OptFunc {
	return func(wid *Widget) {
		key := "class"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), true)
	}
}

func Valign(in AlignType) OptFunc {
	return func(wid *Widget) {
		key := "valign"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), true)
	}
}

func Halign(in AlignType) OptFunc {
	return func(wid *Widget) {
		key := "halign"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), true)
	}
}

func Hexpand(in bool) OptFunc {
	return func(wid *Widget) {
		key := "hexpand"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Vexpand(in bool) OptFunc {
	return func(wid *Widget) {
		key := "vexpand"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Width(in int) OptFunc {
	return func(wid *Widget) {
		key := "width"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Heigth(in int) OptFunc {
	return func(wid *Widget) {
		key := "heigth"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Active(in bool) OptFunc {
	return func(wid *Widget) {
		key := "active"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Visible(in bool) OptFunc {
	return func(wid *Widget) {
		key := "visible"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Tooltip(in string) OptFunc {
	return func(wid *Widget) {
		key := "tooltip"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), true)
	}
}

func PrintStyle(in map[string]string) string {
	var s string

	for k, v := range in {
		s += fmt.Sprintf("%s: %s; ", k, v)
	}

	return s
}

func Style(in map[string]string) OptFunc {
	return func(wid *Widget) {
		key := "style"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(PrintStyle(in)), true)
	}
}

// Orientation function accepts an OrientationType
// it can be VERTICAL or HORIZONTAL
func Orientation(in OrientationType) OptFunc {
	return func(wid *Widget) {
		if wid.Type != BOX && wid.Type != CENTERBOX && wid.Type != SCALE && wid.Type != PROGRESS {
			return
		}
		key := "orientation"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), true)
	}
}

func RoundDigits(in int) OptFunc {
	return func(wid *Widget) {
		if wid.Type != SCALE {
			return
		}
		key := "round-digits"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Flipped(in bool) OptFunc {
	return func(wid *Widget) {
		if wid.Type != SCALE && wid.Type != PROGRESS {
			return
		}
		key := "flipped"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Value(in float64) OptFunc {
	return func(wid *Widget) {
		if wid.Type != SCALE && wid.Type != PROGRESS {
			return
		}
		key := "value"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func DrawValue(in bool) OptFunc {
	return func(wid *Widget) {
		if wid.Type != SCALE {
			return
		}
		key := "draw-value"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func marks(in string) OptFunc {
	// TODO make a more sufficticated form for mark
	return func(wid *Widget) {
		if wid.Type != SCALE {
			return
		}
		key := "marks"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), true)
	}
}

func Min(in float64) OptFunc {
	return func(wid *Widget) {
		if wid.Type != SCALE {
			return
		}
		key := "min"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Max(in float64) OptFunc {
	return func(wid *Widget) {
		if wid.Type != SCALE {
			return
		}
		key := "max"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func TimeOut(in string) OptFunc {
	// TODO make a more suffisticated form
	return func(wid *Widget) {
		if wid.Type != SCALE {
			return
		}
		key := "timeout"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), true)
	}
}

func OnChange(in string) OptFunc {
	return func(wid *Widget) {
		if wid.Type != SCALE && wid.Type != COMBOBOXTEXT {
			return
		}
		key := "onchange"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), true)
	}
}

func Name(in string) OptFunc {
	return func(wid *Widget) {
		if wid.Type != EXPANDER {
			return
		}
		key := "name"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), true)
	}
}

func Expanded(in bool) OptFunc {
	return func(wid *Widget) {
		if wid.Type != EXPANDER {
			return
		}
		key := "expanded"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Text(in string) OptFunc {
	return func(wid *Widget) {
		if wid.Type != LABEL {
			return
		}
		key := "text"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), true)
	}
}

func LimitWidth(in int) OptFunc {
	return func(wid *Widget) {
		if wid.Type != LABEL {
			return
		}
		key := "limit-width"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func ShowTruncated(in bool) OptFunc {
	return func(wid *Widget) {
		if wid.Type != LABEL {
			return
		}
		key := "show-truncated"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Markup(in bool) OptFunc {
	return func(wid *Widget) {
		if wid.Type != LABEL {
			return
		}
		key := "markup"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Wrap(in bool) OptFunc {
	return func(wid *Widget) {
		if wid.Type != LABEL {
			return
		}
		key := "wrap"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Angle(in float64) OptFunc {
	return func(wid *Widget) {
		if wid.Type != LABEL {
			return
		}
		key := "angle"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Xalign(in float64) OptFunc {
	return func(wid *Widget) {
		// TODO normalize numbers more than 1 and less than 0
		if wid.Type != LABEL {
			return
		}
		key := "xalign"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func Yalign(in float64) OptFunc {
	return func(wid *Widget) {
		// TODO normalize numbers more than 1 and less than 0
		if wid.Type != LABEL {
			return
		}
		key := "yalign"
		wid.Attributes[key] = NewAttribute(key, fmt.Sprint(in), false)
	}
}

func AddChildren(children ...*Widget) OptFunc {
	return func(w *Widget) {
		if w.Type != BOX && w.Type != CENTERBOX && w.Type != EXPANDER {
			return
		}
		for _, child := range children {
			w.Children = append(w.Children, *child)
		}
	}
}
