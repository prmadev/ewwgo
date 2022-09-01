package ewwgo_test

import (
	"testing"

	"github.com/amirography/ewwgo"
)

var (
	SimpleBox            = ewwgo.NewBox()
	BoxWithSpacing       = ewwgo.NewBox(ewwgo.Spacing(10))
	BoxWithTwoAttributes = ewwgo.NewBox(ewwgo.Spacing(10), ewwgo.SpaceEvenly(true))
	BoxWithClass         = ewwgo.NewBox(ewwgo.Class("thing"))

	BoxWithChild    = ewwgo.NewBox(ewwgo.Spacing(10)).UpdateChildren(&[]ewwgo.Widget{SimpleBox})
	BoxWithChildren = ewwgo.NewBox(ewwgo.Spacing(10)).UpdateChildren(&[]ewwgo.Widget{SimpleBox, SimpleBox})
)

func TestStringWidget(t *testing.T) {
	t.Parallel()

	table := []struct {
		got  string
		want string
	}{
		{ewwgo.StringWidget(SimpleBox), "(box)"},
		{ewwgo.StringWidget(BoxWithSpacing), "(box :spacing 10)"},
		{ewwgo.StringWidget(BoxWithTwoAttributes), "(box :spacing 10 :space-evenly true)"},
		{ewwgo.StringWidget(BoxWithChild), "(box :spacing 10 (box))"},
		{ewwgo.StringWidget(BoxWithChildren), "(box :spacing 10 (box) (box))"},
	}

	for _, row := range table {
		if row.got != row.want {
			t.Errorf("\nGot\t%s,\nWanted\t%s\n", row.got, row.want)
		}
	}
}
