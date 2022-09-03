package ewwgo

import "fmt"

//
// Box Specific Attributes
//

func Spacing(in int) OptFunc {
	return func(wid *Widget) {
		key := "spacing"
		a := NewAttribute(key, fmt.Sprint(in), false)
		wid.Attributes[key] = a
	}
}

func SpaceEvenly(in bool) OptFunc {
	return func(wid *Widget) {
		key := "space-evenly"
		a := NewAttribute(key, fmt.Sprint(in), false)
		wid.Attributes[key] = a
	}
}
