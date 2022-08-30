package ewwgo

// TODO implement Label type
// type Label struct {
// 	Text          string
// 	LimitWidth    int
// 	ShowTruncated bool
// 	Markup        bool
// 	Wrap          bool
// 	Angle         float64
// 	Xalign        float64
// 	Yalign        float64

// 	*General
// }

// func (m Label) String() string {
// 	var attr []string
// 	attr = append(attr, fmt.Sprintf(":text '%s' ", m.Text))
// 	attr = append(attr, fmt.Sprintf(":limit-width %d ", m.LimitWidth))
// 	// attr = append(attr, fmt.Sprintf(":markup %t", m.Markup))
// 	attr = append(attr, fmt.Sprintf(":wrap %t", m.Wrap))
// 	attr = append(attr, fmt.Sprintf(":angle %f", m.Angle))
// 	attr = append(attr, fmt.Sprintf(":xalign %f", m.Xalign))
// 	attr = append(attr, fmt.Sprintf(":yalign %f", m.Yalign))
// 	attr = append(attr, fmt.Sprintf(":show-truncated %t", m.ShowTruncated))
// 	attr = append(attr, m.General.String()...)

// 	return fmt.Sprintf("(label %s)", stringBuilder(attr))
// }

// func NewLabel() *Label {
// 	g := NewGeneral()

// 	l := &Label{
// 		Text:          "",
// 		LimitWidth:    40,
// 		ShowTruncated: false,
// 		Markup:        false,
// 		Wrap:          false,
// 		Angle:         0,
// 		Xalign:        0.5,
// 		Yalign:        0.5,

// 		General: g,
// 	}

// 	return l
// }
