package ewwgo

// Scale type is a widget of type WidgetDetails
// You should not instantiate a Scale directly
// use NewScale() function
type Scale WidgetDetails

//
// implement as Widget
//

// GetWidget method returns the details of the receiver Scale.
// Mostly used to the purpose of converting to string
func (m *Scale) GetWidget() WidgetDetails {
	return WidgetDetails(*m)
}

// GetAttributes method returns a pointer to attributes so to be changed by the any Option.Apply function
func (m *Scale) GetAttributes() *AttributeSet {
	return m.Attributes
}

//
// Misc
//

// NewScale Function Returns a new widet of type Scale with field of type "Scale"
func NewScale(options ...ScaleOption) *Scale {
	var b Scale
	b.Type = "Scale"
	b.Attributes, _ = NewAttributeSet()

	for _, option := range options {
		option.Apply(&b)
	}

	return &b
}
