package euler

import "testing"

func TestPE0009(t *testing.T) {
	expected := 31875000
	if actual := PE0009(); actual != expected {
		t.Errorf(
			"Expected:%v\nActual:%v",
			expected,
			actual,
		)
	}
}
