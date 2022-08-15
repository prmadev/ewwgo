package ewwgo

import "fmt"

type Widget struct {
	Class   string
	Valign  string
	Halign  string
	Vexpand bool
	Hexpand bool
	Width   int
	Heigth  int
	Active  bool
	Tooltip string
	Visible bool
	Style   string
}

type Box struct {
	Spacing     int
	Orientation string
	SpaceEvenly bool
	Content     []string

	Widget
}

type Progress struct {
	Flipped     bool
	Orientation string
	Value       float64

	Widget
}

type Label struct {
	Text          string
	LimitWidth    int
	ShowTruncated bool
	Markup        bool
	Wrap          bool
	Angle         float64
	Xalign        float64
	Yalign        float64

	Widget
}

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

func (b *Box) Printer() string {

	var c string
	for _, v := range b.Content {
		c = c + v
	}
	return fmt.Sprintf("(box :class '%s' :spacing %d :orientation '%s' :space-evenly %s %s)", b.Class, b.Spacing, b.Orientation, boolToString(b.SpaceEvenly), c)
}

func (b *Label) Printer() string {

	if b.LimitWidth == 0 {
		b.LimitWidth = 30
	}

	class := fmt.Sprintf(":class '%s'", b.Class)
	text := fmt.Sprintf(":text '%s'", b.Text)
	limitWidth := fmt.Sprintf(":limit-width %d", b.LimitWidth)
	markup := ""
	// markup := fmt.Sprintf(":markup %s", boolToString(b.Markup))
	// wrap := fmt.Sprintf(":wrap %s", boolToString(b.Wrap))
	wrap := ""
	angle := fmt.Sprintf(":angle %f", b.Angle)
	xalign := fmt.Sprintf(":xalign %f", b.Xalign)
	yalign := fmt.Sprintf(":yalign %f", b.Yalign)

	return fmt.Sprintf("(label %s %s %s %s %s %s %s %s)", class, text, limitWidth, markup, wrap, angle, xalign, yalign)
}

func (p *Progress) Printer() string {
	class := fmt.Sprintf(":class '%s'", p.Class)
	flipped := fmt.Sprintf(":flipped %s", boolToString(p.Flipped))
	orientation := fmt.Sprintf(":orientation '%s'", p.Orientation)
	value := fmt.Sprintf(":value %f", p.Value)

	return fmt.Sprintf("(progress %s %s %s %s)", class, flipped, orientation, value)
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func (s *Scale) Printer() string {
	class := fmt.Sprintf(":class '%s'", s.Class)
	flipped := fmt.Sprintf(":flipped %s", boolToString(s.Flipped))
	orientation := fmt.Sprintf(":orientation '%s'", s.Orientation)
	value := fmt.Sprintf(":value %f", s.Value)
	marks := fmt.Sprintf(":marks [%s]", s.Marks)
	drawValue := fmt.Sprintf(":draw-value %s", boolToString(s.DrawValue))

	min := fmt.Sprintf(":min %f", s.Min)
	max := fmt.Sprintf(":max %f", s.Max)
	timeout := fmt.Sprintf(":timeout '%s'", s.Timeout)
	onchange := fmt.Sprintf(":onchange '%s'", s.Onchange)

	return fmt.Sprintf("(scale %s %s %s %s %s %s %s %s %s %s)", class, flipped, orientation, value, marks, drawValue, min, max, timeout, onchange)
}
