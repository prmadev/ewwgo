package ewwgo

import "fmt"

type OrientationType string

const (
	VERTICAL   OrientationType = "vertical"
	HORIZONTAL OrientationType = "horizontal"
)

// Orientation function accepts an OrientationType
// it can be VERTICAL or HORIZONTAL
func Orientation(in OrientationType) OptFunc {
	return func(wid *Widget) {
		key := "orientation"
		a := NewAttribute(key, fmt.Sprint(in), true)
		wid.Attributes[key] = a
	}
}
