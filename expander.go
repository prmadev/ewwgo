package ewwgo

func NewExpander(opts ...OptFunc) *Widget {
	var b Widget

	b.Type = "expander"
	b.Attributes = NewAttributeSet()

	for _, opt := range opts {
		opt(&b)
	}

	return &b
}
