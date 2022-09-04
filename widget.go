package ewwgo

import "fmt"

type WidgetType string

const (
	BOX              WidgetType = "box"
	PROGRESS         WidgetType = "progress"
	LABEL            WidgetType = "label"
	SCALE            WidgetType = "scale"
	EXPANDER         WidgetType = "expander"
	COMBOBOXTEXT     WidgetType = "combo-box-text"
	REVEALER         WidgetType = "revealer"
	CHECKBOX         WidgetType = "checkbox"
	COLORBUTTON      WidgetType = "color-button"
	COLORCHOOSER     WidgetType = "color-chooser"
	INPUT            WidgetType = "input"
	BUTTON           WidgetType = "button"
	IMAGE            WidgetType = "image"
	OVERLAY          WidgetType = "overlay"
	CENTERBOX        WidgetType = "centerbox"
	SCROLL           WidgetType = "scroll"
	EVENTBOX         WidgetType = "eventbox"
	LITERAL          WidgetType = "literal"
	CALENDAR         WidgetType = "calendar"
	TRANSFORM        WidgetType = "transform"
	CIRCULARPROGRESS WidgetType = "circularprogress"
	GRAPH            WidgetType = "graph"
)

// WidgetDetails implements dummy info for widgetes
type Widget struct {
	Attributes map[string]Attribute // a map to specific attributes
	Type       WidgetType           // this is the type used to tag the structure in yuck. for example a Type "box" will create a "(box )" string
	Children   []Widget             // the children widgets only some of the widgets have children, they will implement the Parent interface
}

// StringWidget function creates a yuck representation of the given widget and all its subtimes
func StringWidget(m Widget) string {
	var s string

	// TODO Control for empty type
	s += fmt.Sprint(m.Type)

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
