package ewwgo

func NewLabel(opts ...OptFunc) *Widget {
	var b Widget

	b.Type = "label"
	b.Attributes = NewAttributeSet()

	for _, opt := range opts {
		opt(&b)
	}

	return &b
}
