package ewwgo

// Parent interface creates an interface for the Widgets that have children
type Parent interface {
	GetChildren() *[]Widget
	UpdateChildren(*[]Widget)
}

// AddToParent function adds a widget to the Parent widget
func AddToParent(m Widget, in Parent) {
	ws := *in.GetChildren()
	ws = append(ws, m)
	in.UpdateChildren(&ws)
}
