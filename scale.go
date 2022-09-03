package ewwgo

func NewScale(opts ...OptFunc) *Widget {
	var b Widget

	b.Type = "scale"
	b.Attributes = NewAttributeSet()

	for _, opt := range opts {
		opt(&b)
	}

	return &b
}
