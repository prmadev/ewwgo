package ewwgo_test

import (
	"testing"

	"github.com/amirography/ewwgo"
)

func TestNewLabel(t *testing.T) {
	t.Parallel()

	a := ewwgo.NewAttributeSet(ewwgo.NewAttribute("class", "test", true))
	var want ewwgo.Widget = ewwgo.Widget{Type: "label", Attributes: a}
	g := ewwgo.NewLabel(ewwgo.Class("test"))
	got := *g
	if len(want.Attributes) != len(got.Attributes) {
		t.Errorf("want %v, got %v", want, got)
	}

	if len(want.Children) != len(got.Children) {
		t.Errorf("want %v, got %v", want, got)
	}

	if len(want.Type) != len(got.Type) {
		t.Errorf("want %v, got %v", want, got)
	}
}
