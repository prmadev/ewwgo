package ewwgo

import (
	"errors"
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
	k := m.Key

	if m.Wrap {
		k = "'" + k + "'"
	}

	return fmt.Sprintf(":%s %s", m.Key, m.Value)
}

// NewAttribute func returns a pointer to new Attribute.
// it should have a non-nil key and values, other wise it will return an error
func NewAttribute(key, value string, wrap bool) (*Attribute, error) {
	if len(key) == 0 {
		return nil, errors.New("error in creating Attribute: key cannot be empty")
	}

	if len(value) == 0 {
		return nil, errors.New("error in creating Attribute: value cannot be empty")
	}

	return &Attribute{Key: key, Value: value, Wrap: wrap}, nil
}

// AttributeSet is a map to the attributes for a given widget
type AttributeSet map[string]*Attribute

// NewAttributeSet creates a new AttributeSet.
// it will optionally get attributes at creation time
func NewAttributeSet(attributes ...*Attribute) (*AttributeSet, error) {
	m := make(AttributeSet, 0)
	for _, attr := range attributes {
		m[attr.Key] = attr
	}

	return &m, nil
}

// AddAttributes will add one or more pointers to Attribute to an AttributeSet
func AddAttributes(m *AttributeSet, attributes ...*Attribute) {
	as := *m
	for _, attr := range attributes {
		as[attr.Key] = attr
	}
	m = &as
}

// String method will print attributes of the receiver
func (m *AttributeSet) String() string {
	var s string
	for _, v := range *m {
		s += " "
		s += v.String()
	}
	return s
}
