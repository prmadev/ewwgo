package eww

import "fmt"

type Box struct {
	Class       string
	Spacing     int
	Orientation string
	SpaceEvenly bool
	Content     []string
}
type Progress struct {
	Class       string
	Flipped     bool
	Orientation string
	Value       float64
}
type Label struct {
	Class         string
	Text          string
	LimitWidth    int
	ShowTruncated bool
	Markup        bool
	Wrap          bool
	Angle         float64
	Xalign        float64
	Yalign        float64
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
