package ewwgo

import "fmt"

type Box struct {
	Spacing     int
	Orientation string
	SpaceEvenly bool
	Content     []string

	Widget
}

func (b *Box) Printer() string {

	var c string
	for _, v := range b.Content {
		c = c + v
	}
	return fmt.Sprintf("(box :class '%s' :spacing %d :orientation '%s' :space-evenly %t %s)", b.Class, b.Spacing, b.Orientation, b.SpaceEvenly, c)
}
