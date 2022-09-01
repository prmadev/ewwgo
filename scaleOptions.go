package ewwgo

import "fmt"

// ScaleOption interface creates an interface for Scale specific attributes
type ScaleOption interface {
	Apply(*Scale)
}

//
// Scale Specific Attributes
//

// RoundDigits type specifies the the number of decimal digits when showing value
// WARNING for some reason my eww cannot find this option
type RoundDigits int

// Apply method applies the RoundDigits to the given box
func (m RoundDigits) Apply(in *Scale) {
	a, _ := NewAttribute("round-digits", fmt.Sprint(m), false)
	AddAttributes(in.Attributes, a)
}

// TODO	Flipped     bool
// orienters
// TODO	Orientation string
// value taker
// TODO Value       float64
// TODO DrawValue   bool
// TODO Marks       string // disable
// TODO Min         float64
// TODO Max         float64
// TODO Timeout     string // format should be formated as '10s'
// TODO Onchange    string
