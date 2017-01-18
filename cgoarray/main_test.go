package cgoarray

import "testing"

func TestGetArray(t *testing.T) {
	should := []int32{0, 1, 2}
	is, p := GetArray()
	defer FreeArray(p)
	if len(is) != len(should) {
		t.Errorf("Got size '%v', expected '%v'", len(is), len(should))
	}
	for i := range is {
		if is[i] != should[i] {
			t.Errorf("Got value '%v' at index '%v', expected '%v'", is[i], i, should[i])
		}
	}
}
