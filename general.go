package ewwgo

import (
	"fmt"
)

// GeneralValueParams interface implements generic type parameters that acts as input for the attributes that are used by every widget
type GeneralValueParams interface {
	~int | ~string | ~float64 | ~bool
}

// ApplyGeneralBox is a function that is used as a helper for Apply method for the ease of use of general Attributes
func ApplyGeneralBox[P GeneralValueParams](attrs *AttributeSet, key string, param P, wrap bool) {
	a, _ := NewAttribute(key, fmt.Sprint(param), wrap)

	AddAttributes(attrs, a)
}

// Class is a general Attribute of type string
type Class string

// Apply method applies the the Class Attribute to the give widget
func (m Class) Apply(in Widget) {
	ApplyGeneralBox(in.GetAttributes(), "class", m, true)
}

// TODO Document
type Valign string

// TODO Document
func (m Valign) Apply(in Widget) {
	ApplyGeneralBox(in.GetAttributes(), "valign", m, true)
}

// TODO Document
type Halign string

// TODO Document
func (m Halign) Apply(in Widget) {
	ApplyGeneralBox(in.GetAttributes(), "halign", m, true)
}

// TODO Document
type Hexpand bool

// TODO Document
func (m Hexpand) Apply(in Widget) {
	ApplyGeneralBox(in.GetAttributes(), "hexpand", fmt.Sprint(m), false)
}

// TODO Document
type Vexpand bool

// TODO Document
func (m Vexpand) Apply(in Widget) {
	ApplyGeneralBox(in.GetAttributes(), "Vexpand", fmt.Sprint(m), false)
}

// TODO Document
type Width int

// TODO Document
func (m Width) Apply(in Widget) {
	ApplyGeneralBox(in.GetAttributes(), "width", fmt.Sprint(m), false)
}

// TODO Document
type Height int

// TODO Document
func (m Height) Apply(in Widget) {
	ApplyGeneralBox(in.GetAttributes(), "heigth", fmt.Sprint(m), false)
}

// TODO Document
type Active bool

// TODO Document
func (m Active) Apply(in Widget) {
	ApplyGeneralBox(in.GetAttributes(), "Active", fmt.Sprint(m), false)
}

// TODO Document
type Visible bool

// TODO Document
func (m Visible) Apply(in Widget) {
	ApplyGeneralBox(in.GetAttributes(), "Visible", fmt.Sprint(m), false)
}

// TODO Document
type Tooltip string

// TODO Document
func (m Tooltip) Apply(in Widget) {
	ApplyGeneralBox(in.GetAttributes(), "tooltip", m, true)
}

// TODO Document
type Style map[string]string

// TODO Document
func (m Style) String() string {
	var s string

	for k, v := range m {
		s += fmt.Sprintf("%s: %s; ", k, v)
	}

	return s
}

// TODO Document
func (m Style) Apply(in Widget) {
	ApplyGeneralBox(in.GetAttributes(), "style", m.String(), true)
}
