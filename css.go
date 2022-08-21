package ewwgo

import "fmt"

type Css map[string]string

func NewStyle() *Css {
	a := new(Css)
	return a
}

func (m *Css) String() string {
	s := ""

	for k, v := range *m {
		l := fmt.Sprintf("%s: %s;", k, v)
		s = s + l
	}

	return s
}
