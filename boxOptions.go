package ewwgo

import "fmt"

// BoxOption interface creates an interface for Box specific attributes
type BoxOption interface {
	Spacing
	// Apply(*Box)
}

//
// Box Specific Attributes
//

// Spacing type specifies the amount of space between children
type Spacing int

// Apply method applies the spacing to the given box
func (m Spacing) Apply(in *Box) {
	a, _ := NewAttribute("spacing", fmt.Sprint(m), false)
	AddAttributes(in.Attributes, a)
}

// // Orientation type is a string
// // it can be one of these options:
// // "vertical", "horizontal", "v", "h"
// type Orientation string

// // Apply method applies the Orientation to the input box
// func (m Orientation) Apply(in *Box) {
// 	a, _ := NewAttribute("orientation", string(m), false)
// 	AddAttributes(in.Attributes, a)
// }

// SpaceEvenly specifies of the children of the box should have equal space or not
type SpaceEvenly bool

// Apply method applies the SpaceEvenely to the input box
func (m SpaceEvenly) Apply(in *Box) {
	a, _ := NewAttribute("space-evenly", fmt.Sprint(m), false)
	AddAttributes(in.Attributes, a)
}
