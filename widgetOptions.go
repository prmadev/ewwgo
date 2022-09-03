package ewwgo

type OptFunc func(*Widget)

func (m *Widget) Option(opts ...OptFunc) *Widget {
	for _, opt := range opts {
		opt(m)
	}
	return m
}
