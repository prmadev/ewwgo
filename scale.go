package ewwgo

import "fmt"

type Scale struct {
	Flipped     bool
	Orientation string
	Value       float64
	DrawValue   bool
	Marks       string
	Min         float64
	Max         float64
	Timeout    	string 
	Onchange    string

	Widget
}

func (s *Scale) Printer() string {
	class := fmt.Sprintf(":class '%s'", s.Class)
	flipped := fmt.Sprintf(":flipped %t", s.Flipped)
	orientation := fmt.Sprintf(":orientation '%s'", s.Orientation)
	value := fmt.Sprintf(":value %f", s.Value)
	marks := fmt.Sprintf(":marks '%s'", s.Marks)
	drawValue := fmt.Sprintf(":draw-value %t", s.DrawValue)

	min := fmt.Sprintf(":min %f", s.Min)
	max := fmt.Sprintf(":max %f", s.Max)
	timeout := fmt.Sprintf(":timeout %s", s.Timeout)
	onchange := fmt.Sprintf(":onchange '%s'", s.Onchange)

	return fmt.Sprintf("(scale %s %s %s %s %s %s %s %s %s %s)", class, flipped, orientation, value, marks, drawValue, min, max, timeout, onchange)
}
