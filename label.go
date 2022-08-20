package ewwgo

import (
	"errors"
	"fmt"
)

type Label struct {
	Text          string
	LimitWidth    int
	ShowTruncated bool
	Markup        bool
	Wrap          bool
	Angle         float64
	Xalign        float64
	Yalign        float64

	*General
}

func (m Label) String() string {
	if m.LimitWidth == 0 {
		m.LimitWidth = 30
	}

	class := fmt.Sprintf(":class '%s' ", m.Class)
	text := fmt.Sprintf(":text '%s' ", m.Text)
	limitWidth := fmt.Sprintf(":limit-width %d ", m.LimitWidth)
	// markup := ""
	// markup := fmt.Sprintf(":markup %s", boolToString(b.Markup))
	// wrap := fmt.Sprintf(":wrap %s", boolToString(b.Wrap))
	// wrap := ""
	angle := fmt.Sprintf(":angle %f ", m.Angle)
	xalign := fmt.Sprintf(":xalign %f ", m.Xalign)
	yalign := fmt.Sprintf(":yalign %f ", m.Yalign)

	return fmt.Sprintf("(label %s %s %s %s %s %s)", class, text, limitWidth, angle, xalign, yalign)
}

func NewLabel() *Label {
	g := NewGeneral()

	l := &Label{
		Text:          "",
		LimitWidth:    40,
		ShowTruncated: false,
		Markup:        false,
		Wrap:          false,
		Angle:         0,
		Xalign:        0.5,
		Yalign:        0.5,

		General: g,
	}

	return l
}

func (m *Label) SetText(s string) error {
	m.Text = s

	return nil
}

func (m *Label) SetLimitWidth(d int) error {
	m.LimitWidth = d

	return nil
}

func (m *Label) SetShowTruncated(t bool) error {
	m.ShowTruncated = t

	return nil
}

func (m *Label) SetMarkup(t bool) error {
	m.Markup = t

	return nil
}

func (m *Label) SetWrap(t bool) error {
	m.Wrap = t

	return nil
}

func (m *Label) SetAngle(f float64) error {
	m.Angle = f

	return nil
}

func (m *Label) SetXAlign(f float64) error {
	if f > 1 {
		return errors.New("err: cannot be more than 1")
	}
	m.Xalign = f

	return nil
}

func (m *Label) SetYAlign(f float64) error {
	if f > 1 {
		return errors.New("err: cannot be more than 1")
	}
	m.Yalign = f

	return nil
}

func (m *Label) SetGeneral(g *General) error {
	m.General = g

	return nil
}
