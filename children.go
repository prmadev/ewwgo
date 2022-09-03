package ewwgo

func AddChildren(children ...*Widget) OptFunc {
	return func(w *Widget) {
		for _, child := range children {
			w.Children = append(w.Children, *child)
		}
	}
}
