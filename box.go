package ewwgo

// NewBox Function Returns a new widet of type box with field of type "Box"
func NewBox(opts ...OptFunc) *Widget {
	var b Widget

	b.Type = "box"
	b.Attributes = NewAttributeSet()

	for _, opt := range opts {
		opt(&b)
	}

	return &b
}
