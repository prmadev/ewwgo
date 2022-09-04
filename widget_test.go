package ewwgo_test

import (
	"testing"

	"github.com/amirography/ewwgo"
)

func TestStringWidget(t *testing.T) {
	t.Parallel()
	var (
		SimpleBox              = ewwgo.NewBox()
		BoxWithSpacing         = ewwgo.NewBox(ewwgo.Spacing(10))
		BoxWithThreeAttributes = ewwgo.NewBox(
			ewwgo.SpaceEvenly(true),
			ewwgo.Spacing(10),
			ewwgo.Orientation(ewwgo.VERTICAL))
		BoxWithChild = ewwgo.NewBox(
			ewwgo.Spacing(10),
			ewwgo.AddChildren(SimpleBox))
		BoxWithChildren = ewwgo.NewBox(ewwgo.Spacing(10),
			ewwgo.AddChildren(SimpleBox, SimpleBox))
	)

	table := []struct {
		got  string
		want string
	}{
		{ewwgo.StringWidget(*SimpleBox), "(box)"},
		{ewwgo.StringWidget(*BoxWithSpacing), "(box :spacing 10)"},
		{ewwgo.StringWidget(*BoxWithThreeAttributes), "(box :space-evenly true :spacing 10 :orientation vertical)"},
		{ewwgo.StringWidget(*BoxWithChild), "(box :spacing 10 (box))"},
		{ewwgo.StringWidget(*BoxWithChildren), "(box :spacing 10 (box) (box))"},
	}
	for _, row := range table {
		if row.got != row.want {
			t.Errorf("\nGot\t%s,\nWanted\t%s\n", row.got, row.want)
		}
	}
}
