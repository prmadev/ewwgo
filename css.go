package ewwgo

import "fmt"

type Css map[string]string

func NewStyle() Css {
	a := make(Css, 0)

	return a
}

func (m *Css) String() string {
	s := " "

	for k, v := range *m {
		s += fmt.Sprintf("%s: %s; ", k, v)
	}

	return s
}

