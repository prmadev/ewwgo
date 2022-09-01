package ewwgo

// Label type is a widget of type WidgetDetails
// You should not instantiate a Label directly
// use NewLabel() function
type Label WidgetDetails

//
// implement as Widget
//

// GetWidget method returns the details of the receiver label.
// Mostly used to the purpose of converting to string
func (m *Label) GetWidget() WidgetDetails {
	return WidgetDetails(*m)
}

// GetAttributes method returns a pointer to attributes so to be changed by the any Option.Apply function
func (m *Label) GetAttributes() *AttributeSet {
	return m.Attributes
}

// GetOrientorAttributes method returns a pointer to attributes so to be changed by the any Option.Apply function
func (m *Label) GetOrientorAttributes() *AttributeSet {
	return m.Attributes
}

//
// Misc
//

// NewLabel Function Returns a new widet of type label with field of type "Label"
func NewLabel(options ...LabelOption) *Label {
	var b Label
	b.Type = "label"
	b.Attributes, _ = NewAttributeSet()

	for _, option := range options {
		option.Apply(&b)
	}

	return &b
}
