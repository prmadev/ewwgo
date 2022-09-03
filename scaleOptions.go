package ewwgo

// import "fmt"

// // Scale type is a widget of type WidgetDetails
// // You should not instantiate a Scale directly
// // use NewScale() function
// type Scale WidgetDetails

// //
// // implement as Widget
// //

// // GetWidget method returns the details of the receiver Scale.
// // Mostly used to the purpose of converting to string
// func (m *Scale) GetWidget() WidgetDetails {
// 	return WidgetDetails(*m)
// }

// // GetAttributes method returns a pointer to attributes so to be changed by the any Option.Apply function
// func (m *Scale) GetAttributes() *AttributeSet {
// 	return m.Attributes
// }

// //
// // Misc
// //

// // NewScale Function Returns a new widet of type Scale with field of type "Scale"
// func NewScale(options ...ScaleOption) *Scale {
// 	var b Scale
// 	b.Type = "Scale"
// 	b.Attributes, _ = NewAttributeSet()

// 	for _, option := range options {
// 		option.Apply(&b)
// 	}

// 	return &b
// }
// // ScaleOption interface creates an interface for Scale specific attributes
// type ScaleOption interface {
// 	Apply(*Scale)
// }

// //
// // Scale Specific Attributes
// //

// // RoundDigits type specifies the the number of decimal digits when showing value
// // WARNING for some reason my eww cannot find this option
// type RoundDigits int

// // Apply method applies the RoundDigits to the given box
// func (m RoundDigits) Apply(in *Scale) {
// 	a, _ := NewAttribute("round-digits", fmt.Sprint(m), false)
// 	AddAttributes(in.Attributes, a)
// }

// // TODO	Flipped     bool
// // orienters
// // TODO	Orientation string
// // value taker
// // TODO Value       float64
// // TODO DrawValue   bool
// // TODO Marks       string // disable
// // TODO Min         float64
// // TODO Max         float64
// // TODO Timeout     string // format should be formated as '10s'
// // TODO Onchange    string
