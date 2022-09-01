package ewwgo_test

import (
	"testing"

	"github.com/amirography/ewwgo"
)

func TestNewBox(t *testing.T) {
	t.Parallel()

	a, _ := ewwgo.NewAttributeSet()
	var want ewwgo.Box = ewwgo.Box{Type: "box", Attributes: a}
	g := ewwgo.NewBox()
	got := *g
	if len(*want.Attributes) != len(*got.Attributes) {
		t.Errorf("want %v, got %v", want, got)
	}

	if len(want.Children) != len(got.Children) {
		t.Errorf("want %v, got %v", want, got)
	}

	if len(want.Type) != len(got.Type) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestGetAttribute(t *testing.T) {
	t.Parallel()

	atr, _ := ewwgo.NewAttributeSet()
	b := ewwgo.Box{Type: "box", Attributes: atr}

	want := b.Attributes
	got := b.GetAttributes()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestGetWidget(t *testing.T) {
	t.Parallel()

	atr, _ := ewwgo.NewAttributeSet()
	b := ewwgo.Box{Type: "box", Attributes: atr}

	want := b.Attributes
	got := b.GetAttributes()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
