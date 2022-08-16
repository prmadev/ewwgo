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
	Timeout     string
	Onchange    string

	Widget
}

func (s Scale) Printer() string {
	var attr []string
	attr = append(attr, fmt.Sprintf(":flipped %t ", s.Flipped))
	attr = append(attr, fmt.Sprintf(":orientation '%s' ", s.Orientation))
	attr = append(attr, fmt.Sprintf(":value %f ", s.Value))
	attr = append(attr, fmt.Sprintf(":marks '%s' ", s.Marks))
	attr = append(attr, fmt.Sprintf(":draw-value %t ", s.DrawValue))
	// attr = append(attr, fmt.Sprintf(":min %f ", s.Min))
	// attr = append(attr, fmt.Sprintf(":max %f ", s.Max))
	// attr = append(attr, fmt.Sprintf(":timeout %s ", s.Timeout))
	// attr = append(attr, fmt.Sprintf(":onchange '%s' ", s.Onchange))
	attr = append(attr, s.Widget.Printer()...)

	str := stringBuilder(attr)

	return fmt.Sprintf("(scale %s)", str)
}
