package ewwgo_test

import (
	"testing"

	"github.com/amirography/ewwgo"
)

func TestNewBox(t *testing.T) {
	t.Parallel()

	a := ewwgo.NewAttributeSet(ewwgo.NewAttribute("class", "test", true))
	var want ewwgo.Widget = ewwgo.Widget{Type: "box", Attributes: a}
	g := ewwgo.NewBox(ewwgo.Class("test"))
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
