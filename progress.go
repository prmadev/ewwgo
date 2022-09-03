package ewwgo

func NewProgress(opts ...OptFunc) *Widget {
	var b Widget

	b.Type = "progress"
	b.Attributes = NewAttributeSet()

	for _, opt := range opts {
		opt(&b)
	}

	return &b
}
