package ewwgo

// Widget interface applies to every widget type
// it should be only implemented where the reciever is WidgetDetails
// type Widget interface {
// Option(OptFunc)
// GetWidget() WidgetDetails
// }

// WidgetDetails implements dummy info for widgetes
type Widget struct {
	Attributes map[string]Attribute // a map to specific attributes
	Type       string               // this is the type used to tag the structure in yuck. for example a Type "box" will create a "(box )" string
	Children   []Widget             // the children widgets only some of the widgets have children, they will implement the Parent interface
}

// StringWidget function creates a yuck representation of the given widget and all its subtimes
func StringWidget(m Widget) string {
	var s string

	// TODO Control for empty type
	s += m.Type

	for _, v := range m.Attributes {
		s += " "
		s += v.String()
	}

	for _, w := range m.Children {
		s += " "
		s += StringWidget(w)
	}

	s = "(" + s + ")"

	return s
}
