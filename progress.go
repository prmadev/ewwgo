package ewwgo

import "fmt"

type Progress struct {
	Flipped     bool
	Orientation string
	Value       float64

	Widget
}

func (p *Progress) Printer() string {
	class := fmt.Sprintf(":class '%s'", p.Class)
	flipped := fmt.Sprintf(":flipped %t", p.Flipped)
	orientation := fmt.Sprintf(":orientation '%s'", p.Orientation)
	value := fmt.Sprintf(":value %f", p.Value)

	return fmt.Sprintf("(progress %s %s %s %s)", class, flipped, orientation, value)
}
