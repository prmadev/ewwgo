package ewwgo

import "fmt"

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

func (b Label) Printer() string {

	if b.LimitWidth == 0 {
		b.LimitWidth = 30
	}

	class := fmt.Sprintf(":class '%s' ", b.Class)
	text := fmt.Sprintf(":text '%s' ", b.Text)
	limitWidth := fmt.Sprintf(":limit-width %d ", b.LimitWidth)
	// markup := ""
	// markup := fmt.Sprintf(":markup %s", boolToString(b.Markup))
	// wrap := fmt.Sprintf(":wrap %s", boolToString(b.Wrap))
	// wrap := ""
	angle := fmt.Sprintf(":angle %f ", b.Angle)
	xalign := fmt.Sprintf(":xalign %f ", b.Xalign)
	yalign := fmt.Sprintf(":yalign %f ", b.Yalign)

	return fmt.Sprintf("(label %s %s %s %s %s %s)", class, text, limitWidth,  angle, xalign, yalign)
}
