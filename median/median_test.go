package median

import "testing"

var (
	tests = []struct {
		name    string
		list    []float64
		medians []float64
	}{
		{
			"41532",
			[]float64{4, 1, 5, 3, 2},
			[]float64{3},
		},
		{
			"1981",
			[]float64{1, 9, 8, 1},
			[]float64{4.5, 1},
		},
		{
			"-1-2",
			[]float64{-1, -2},
			[]float64{-1.5, -2},
		},
		{
			"111",
			[]float64{1, 1, 1},
			[]float64{1},
		},
		{
			"1",
			[]float64{1},
			[]float64{1},
		},
		{
			"empty",
			[]float64{},
			[]float64{0},
		},
	}
	testFns = []struct {
		name string
		fn   func([]float64) float64
	}{
		{
			"median",
			median,
		},
		{
			"median2",
			median2,
		},
	}
)

func TestMedian(t *testing.T) {
	for _, fn := range testFns {
		t.Run(fn.name, func(t *testing.T) {
			for _, test := range tests {
				t.Run(test.name, func(t *testing.T) {
					med := fn.fn(test.list)
					found := false
					for _, m := range test.medians {
						if med == m {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("is %v, want one of %v", med, test.medians)
					}
				})
			}
		})
	}
}
