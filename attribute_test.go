package ewwgo_test

import (
	"testing"

	"github.com/amirography/ewwgo"
)

func TestNewAttribute(t *testing.T) {
	testCases := []struct {
		key   string
		value string
		wrap  bool
		want  ewwgo.Attribute
	}{
		{"", "test", false, ewwgo.Attribute{}},
		{"test", "", false, ewwgo.Attribute{}},
	}
	for _, ts := range testCases {
		got := ewwgo.NewAttribute(ts.key, ts.value, ts.wrap)
		if got.Key != ts.want.Key || got.Value != ts.want.Value || got.Wrap != ts.want.Wrap {
			t.Errorf("\nwanted:\t%v\ngot:\t%v\n", ts.want, got)
		}
	}
}
