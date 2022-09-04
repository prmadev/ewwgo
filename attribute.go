package ewwgo

import (
	"fmt"
)

// Attribute type creates an standard attribute option
// it will end up looking like this:
// :key value
// or if wrap is true:
// :key 'value'
type Attribute struct {
	Key   string
	Value string
	Wrap  bool
}

// String method creates a string representation of receiver attribute
func (m *Attribute) String() string {
	v := m.Value

	if m.Wrap {
		v = "'" + v + "'"
	}

	return fmt.Sprintf(":%s %s", m.Key, v)
}

// NewAttribute func returns a pointer to new Attribute.
// it should have a non-nil key and values, other wise it will return an error
func NewAttribute(key, value string, wrap bool) Attribute {
	if len(key) == 0 {
		return Attribute{}
		// , errors.New("error in creating Attribute: key cannot be empty")
	}

	if len(value) == 0 {
		return Attribute{}
		// , errors.New("error in creating Attribute: value cannot be empty")
	}

	return Attribute{Key: key, Value: value, Wrap: wrap}
	// , nil
}

// NewAttributeSet creates a new AttributeSet.
// it will optionally get attributes at creation time
func NewAttributeSet(attributes ...Attribute) map[string]Attribute {
	m := make(map[string]Attribute, 0)
	for _, attr := range attributes {
		m[attr.Key] = attr
	}

	return m
}
