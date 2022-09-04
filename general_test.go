package ewwgo_test

import (
	"fmt"
	"testing"

	"github.com/amirography/ewwgo"
)

func TestAttrs(t *testing.T) {
	t.Parallel()
	type testCase struct {
		of      ewwgo.OptFunc
		input   any
		atrName string
		on      ewwgo.Widget
	}
	tcs := []testCase{
		{of: ewwgo.Class("test"), input: "test", atrName: "class", on: *ewwgo.NewBox()},
		{of: ewwgo.Spacing(2), input: 2, atrName: "spacing", on: *ewwgo.NewBox()},
		{of: ewwgo.Valign(ewwgo.BASELINE), input: "baseline", atrName: "valign", on: *ewwgo.NewBox()},
		{of: ewwgo.Halign("center"), input: ewwgo.CENTER, atrName: "halign", on: *ewwgo.NewProgress()},
		{of: ewwgo.Vexpand(true), input: true, atrName: "vexpand", on: *ewwgo.NewLabel()},
		{of: ewwgo.Hexpand(false), input: false, atrName: "hexpand", on: *ewwgo.NewScale()},
		{of: ewwgo.Width(10), input: "baseline", atrName: "width", on: *ewwgo.NewExpander()},
		{of: ewwgo.Heigth(40), input: "baseline", atrName: "heigth", on: *ewwgo.NewBox()},
		{of: ewwgo.Visible(true), input: "baseline", atrName: "visible", on: *ewwgo.NewBox()},
		{of: ewwgo.Active(false), input: "baseline", atrName: "active", on: *ewwgo.NewBox()},
		{of: ewwgo.Tooltip("sometext"), input: "baseline", atrName: "tooltip", on: *ewwgo.NewBox()},
		{of: ewwgo.Style(map[string]string{"background": "#ffffff"}), input: map[string]string{"background": "#ffffff"}, atrName: "style", on: *ewwgo.NewBox()},
	}
	for _, tc := range tcs {
		testerfunc(t, tc.of, tc.input, tc.atrName, tc.on)
	}
}

func testerfunc(t *testing.T, of ewwgo.OptFunc, input any, atrName string, on ewwgo.Widget) {
	got, want := on, on
	got.Option(of)
	want.Attributes = ewwgo.NewAttributeSet(ewwgo.NewAttribute(atrName, fmt.Sprint(input), true))

	if len(want.Attributes) != len(got.Attributes) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestCheck(t *testing.T) {
	t.Parallel()
	type testCase struct {
		of      ewwgo.OptFunc
		input   any
		atrName string
		on      ewwgo.Widget
		wanted  ewwgo.Widget
	}
	tcs := []testCase{
		{of: ewwgo.Spacing(2), input: 2, atrName: "spacing", on: *ewwgo.NewExpander(), wanted: *ewwgo.NewExpander()},
	}
	for _, tc := range tcs {
		checktesterfunc(t, tc.of, tc.input, tc.atrName, tc.on, tc.wanted)
	}
}

func checktesterfunc(t *testing.T, of ewwgo.OptFunc, input any, atrName string, got ewwgo.Widget, want ewwgo.Widget) {
	got.Option(of)

	if len(want.Attributes) != len(got.Attributes) {
		t.Errorf("want %v, got %v", want, got)
	}
}
