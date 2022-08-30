package ewwgo

// Widget interface applies to every widget type
// it should be only implemented where the reciever is WidgetDetails
type Widget interface {
	GetWidget() WidgetDetails
	GetAttributes() *AttributeSet
}

// WidgetDetails implements dummy info for widgetes
type WidgetDetails struct {
	Name       string        // the user given name, it is optional
	Attributes *AttributeSet // a pointer to attributeset
	Type       string        // this is the type used to tag the structure in yuck. for example a Type "box" will create a "(box )" string
	Children   []Widget      // the children widgets only some of the widgets have children, they will implement the Parent interface
}

// StringWidget function creates a yuck representation of the given widget and all its subtimes
func StringWidget(m Widget) string {
	var s string

	s += m.GetWidget().Type + " "
	s += m.GetWidget().Attributes.String()

	for _, w := range m.GetWidget().Children {

		s += " "
		s += StringWidget(w)
	}

	s += "(" + s + ")"

	return s
}
