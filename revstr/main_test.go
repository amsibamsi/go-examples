package revstr

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		original, reversed string
	}{
		{"", ""},
		{"x", "x"},
		{"ab", "ba"},
		{"my string", "gnirts ym"},
	}
	for _, test := range tests {
		t.Run(test.original, func(t *testing.T) {
			r := ReverseString(test.original)
			if r != test.reversed {
				t.Errorf("got '%s', expected '%s'", r, test.reversed)
			}
		})
	}
}
