package ewwgo

// // Label type is a widget of type WidgetDetails
// // You should not instantiate a Label directly
// // use NewLabel() function
// type Label WidgetDetails

// //
// // implement as Widget
// //

// // GetWidget method returns the details of the receiver label.
// // Mostly used to the purpose of converting to string
// func (m *Label) GetWidget() WidgetDetails {
// 	return WidgetDetails(*m)
// }

// // GetAttributes method returns a pointer to attributes so to be changed by the any Option.Apply function
// func (m *Label) GetAttributes() *AttributeSet {
// 	return m.Attributes
// }

// // GetOrientorAttributes method returns a pointer to attributes so to be changed by the any Option.Apply function
// func (m *Label) GetOrientorAttributes() *AttributeSet {
// 	return m.Attributes
// }

// //
// // Misc
// //

// // NewLabel Function Returns a new widet of type label with field of type "Label"
// func NewLabel(options ...LabelOption) *Label {
// 	var b Label
// 	b.Type = "label"
// 	b.Attributes, _ = NewAttributeSet()

// 	for _, option := range options {
// 		option.Apply(&b)
// 	}

// 	return &b
// }
// import "fmt"

// // LabelOption interface creates an interface for Label specific attributes
// type LabelOption interface {
// 	Apply(*Label)
// }

// //
// // Label Specific Attributes
// //

// // Text type specifies the Text that is shown in the label
// type Text string

// // Apply method applies the Text to the given Label
// func (m Text) Apply(in *Label) {
// 	a, _ := NewAttribute("text", fmt.Sprint(m), true)
// 	AddAttributes(in.Attributes, a)
// }

// // LimitWidth type limits the number of characters in
// // the label that is shown
// type LimitWidth int

// // Apply method applies the LimitWidth to the given Label
// func (m LimitWidth) Apply(in *Label) {
// 	a, _ := NewAttribute("limit-width", fmt.Sprint(m), false)
// 	AddAttributes(in.Attributes, a)
// }

// // ShowTruncated type specifies if there should be a "..." character when characters are truncated
// type ShowTruncated bool

// // Apply method applies the ShowTruncated to the given Label
// func (m ShowTruncated) Apply(in *Label) {
// 	a, _ := NewAttribute("show-truncated", fmt.Sprint(m), false)
// 	AddAttributes(in.Attributes, a)
// }

// // Markup type specifies if pango markup should be enabled
// // WARNING this option breaks my configuration in my system
// type Markup bool

// // Apply method applies the Markup to the given Label
// func (m Markup) Apply(in *Label) {
// 	a, _ := NewAttribute("markup", fmt.Sprint(m), false)
// 	AddAttributes(in.Attributes, a)
// }

// // Wrap type specifies if wrapping of the label should be enabled
// type Wrap bool

// // Apply method applies the Wrap to the given Label
// func (m Wrap) Apply(in *Label) {
// 	a, _ := NewAttribute("wrap", fmt.Sprint(m), false)
// 	AddAttributes(in.Attributes, a)
// }

// // Angle type specifies the angle at which the should be showing
// type Angle float64

// // Apply method applies the Angle to the given Label
// func (m Angle) Apply(in *Label) {
// 	a, _ := NewAttribute("angle", fmt.Sprint(m), false)
// 	AddAttributes(in.Attributes, a)
// }

// // Xalign type specifies the position of text, in x axis
// // it starts from 0.0 to 1.0
// //  for example: the center is 0.5
// type Xalign float64

// // Apply method applies the Xalign to the given Label
// func (m Xalign) Apply(in *Label) {
// 	a, _ := NewAttribute("xalign", fmt.Sprint(m), false)
// 	AddAttributes(in.Attributes, a)
// }

// // Yalign type specifies the position of text, in y axis
// // it starts from 0.0 to 1.0
// //  for example: the center is 0.5
// type Yalign float64

// // Apply method applies the Yalign to the given Label
// func (m Yalign) Apply(in *Label) {
// 	a, _ := NewAttribute("yalign", fmt.Sprint(m), false)
// 	AddAttributes(in.Attributes, a)
// }
