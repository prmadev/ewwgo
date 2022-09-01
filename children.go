package ewwgo

// Parent interface creates an interface for the Widgets that have children
type Parent interface {
	GetChildren() *[]Widget
	UpdateChildren(*[]Widget) Parent
}

// AddToParent function adds a widget to the Parent widget
func AddToParent(m Widget, in Parent) {
	ws := *in.GetChildren()
	ws = append(ws, m)
	in.UpdateChildren(&ws)
}

// Orientor interface creates an interface for the Widgets that have children
type Orientor interface {
	GetOrientorAttributes() *AttributeSet
	UpdateOrientorAttributes(*AttributeSet) *AttributeSet
}

// Orientation type is a string
// it can be one of these options:
// "vertical", "horizontal", "v", "h"
type Orientation string

// Apply method applies the Orientation to the input box
func (m Orientation) Apply(in Orientor) {
	a, _ := NewAttribute("orientation", string(m), false)
	AddAttributes(in.GetOrientorAttributes(), a)
}
