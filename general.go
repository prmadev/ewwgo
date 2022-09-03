package ewwgo

import "fmt"

type AlignType string

const (
	FILL     AlignType = "fill"
	BASELINE AlignType = "baseline"
	CENTER   AlignType = "center"
	START    AlignType = "start"
	END      AlignType = "end"
)

func Class(in string) OptFunc {
	return func(wid *Widget) {
		key := "class"
		a := NewAttribute(key, fmt.Sprint(in), true)
		wid.Attributes[key] = a
	}
}

func Valign(in AlignType) OptFunc {
	return func(wid *Widget) {
		key := "Valign"
		a := NewAttribute(key, fmt.Sprint(in), true)
		wid.Attributes[key] = a
	}
}

func Halign(in AlignType) OptFunc {
	return func(wid *Widget) {
		key := "Halign"
		a := NewAttribute(key, fmt.Sprint(in), true)
		wid.Attributes[key] = a
	}
}

func Hexpand(in bool) OptFunc {
	return func(wid *Widget) {
		key := "Hexpand"
		a := NewAttribute(key, fmt.Sprint(in), false)
		wid.Attributes[key] = a
	}
}

func Vexpand(in bool) OptFunc {
	return func(wid *Widget) {
		key := "Vexpand"
		a := NewAttribute(key, fmt.Sprint(in), false)
		wid.Attributes[key] = a
	}
}

func Width(in int) OptFunc {
	return func(wid *Widget) {
		key := "width"
		a := NewAttribute(key, fmt.Sprint(in), false)
		wid.Attributes[key] = a
	}
}

func Heigth(in int) OptFunc {
	return func(wid *Widget) {
		key := "heigth"
		a := NewAttribute(key, fmt.Sprint(in), false)
		wid.Attributes[key] = a
	}
}

func Active(in bool) OptFunc {
	return func(wid *Widget) {
		key := "active"
		a := NewAttribute(key, fmt.Sprint(in), false)
		wid.Attributes[key] = a
	}
}

func Visible(in bool) OptFunc {
	return func(wid *Widget) {
		key := "visible"
		a := NewAttribute(key, fmt.Sprint(in), false)
		wid.Attributes[key] = a
	}
}

func Tooltip(in string) OptFunc {
	return func(wid *Widget) {
		key := "tooltip"
		a := NewAttribute(key, fmt.Sprint(in), true)
		wid.Attributes[key] = a
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
		a := NewAttribute(key, fmt.Sprint(PrintStyle(in)), true)
		wid.Attributes[key] = a
	}
}
