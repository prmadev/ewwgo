package ewwgo

func NewLabel(opts ...OptFunc) *Widget {
	var b Widget

	b.Type = LABEL
	b.Attributes = NewAttributeSet()

	for _, opt := range opts {
		opt(&b)
	}

	return &b
}

func NewScale(opts ...OptFunc) *Widget {
	var b Widget

	b.Type = SCALE
	b.Attributes = NewAttributeSet()

	for _, opt := range opts {
		opt(&b)
	}

	return &b
}

func NewProgress(opts ...OptFunc) *Widget {
	var b Widget

	b.Type = PROGRESS
	b.Attributes = NewAttributeSet()

	for _, opt := range opts {
		opt(&b)
	}

	return &b
}

// NewBox Function Returns a new widet of type box with field of type "Box"
func NewBox(opts ...OptFunc) *Widget {
	var b Widget

	b.Type = BOX
	b.Attributes = NewAttributeSet()

	for _, opt := range opts {
		opt(&b)
	}

	return &b
}

func NewExpander(opts ...OptFunc) *Widget {
	var b Widget

	b.Type = EXPANDER
	b.Attributes = NewAttributeSet()

	for _, opt := range opts {
		opt(&b)
	}

	return &b
}
