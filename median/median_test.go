package median

import "testing"

var (
	tests = []struct {
		name   string
		list   []float64
		median float64
	}{
		{
			"1to5",
			[]float64{4, 1, 5, 3, 2},
			3,
		},
		{
			"1189",
			[]float64{1, 9, 8, 1},
			4.5,
		},
		{
			"-1-2",
			[]float64{-2, -1},
			-1.5,
		},
		{
			"single",
			[]float64{1},
			1,
		},
		{
			"empty",
			[]float64{},
			0,
		},
	}
)

func TestMedian(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := median(test.list)
			if m != test.median {
				t.Errorf("median(%v)=%v, want %v", test.list, m, test.median)
			}
		})
	}
}
