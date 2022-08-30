package ewwgo

// Box type is a widget of type WidgetDetails
// You should not instantiate a Box directly
// use NewBox() function
type Box WidgetDetails

//
// implement as Widget
//

// GetWidget method returns the details of the receiver box.
// Mostly used to the purpose of converting to string
func (m *Box) GetWidget() WidgetDetails {
	return WidgetDetails(*m)
}

// GetAttributes method returns a pointer to attributes so to be changed by the any Option.Apply function
func (m *Box) GetAttributes() *AttributeSet {
	return m.Attributes
}

//
// implemet as Parent
//

// GetChildren method recieves a pointer to the slice of children widgets
func (m *Box) GetChildren() *[]Widget {
	return &m.Children
}

// UpdateChildren method updates the reciever Children field with the input widgets
func (m *Box) UpdateChildren(in *[]Widget) {
	m.Children = *in
}

//
// Misc
//

// NewBox Function Returns a new widet of type box with field of type "Box"
func NewBox(options ...BoxOption) *Box {
	var b Box
	b.Type = "box"
	b.Attributes, _ = NewAttributeSet()

	for _, option := range options {
		option.Apply(&b)
	}

	return &b
}
