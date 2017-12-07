package euler

import "testing"

func TestPE0023(t *testing.T) {
	if actual := PE0023(); actual != 4179871 {
		t.Errorf("Expected:%v\nActual:%v", 4179871, actual)
	}
}
