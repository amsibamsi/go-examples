package median

import (
	"math/rand"
	"strconv"
	"testing"
)

var (
	medianFunctions = []struct {
		name string
		fn   func([]float64) float64
	}{
		{
			"median-sort",
			medianSort,
		},
		{
			"median-kselect",
			medianKSelect,
		},
	}
)

var (
	tests = []struct {
		name    string
		list    []float64
		medians []float64
	}{
		{
			"12345",
			[]float64{1, 2, 3, 4, 5},
			[]float64{3},
		},
		{
			"54321",
			[]float64{5, 4, 3, 2, 1},
			[]float64{3},
		},
		{
			"41532",
			[]float64{4, 1, 5, 3, 2},
			[]float64{3},
		},
		{
			"11123",
			[]float64{1, 1, 1, 2, 3},
			[]float64{1},
		},
		{
			"11122",
			[]float64{1, 1, 1, 2, 2},
			[]float64{1},
		},
		{
			"6879",
			[]float64{6, 8, 7, 9},
			[]float64{7.5, 7},
		},
		{
			"9766",
			[]float64{9, 7, 6, 6},
			[]float64{6.5, 6},
		},
		{
			"6766",
			[]float64{6, 7, 6, 6},
			[]float64{6},
		},
		{
			"-1-2",
			[]float64{-1, -2},
			[]float64{-1.5, -1, -2},
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
		{
			"big",
			[]float64{345, 6, 8, 3, 4, 5, 7, 8, 456, 34, 4, 56, 6, 54, 3, 45, 5, 6, 7, 8, 00, 8, 58, 455, 4, 43, 2, 4, 45, 65, 658, 4, 8, 56, 45, 526, 37, 658, 59, 87, 70, 87, 909999, 865.4545, 8.346435, 345.66, 34.65, 54, 7, 45, 3, 435, 346, 45, 56, 869, 879, 87, 49983, 832, 38, 83, 3, 838368, 76476, 4574567, 457436573665, 4325425432, 567657, 6587865456, 73456, 56, 436, 45, 76537, 865, 73, 5654, 6, 436},
			[]float64{56},
		},
	}
)

func TestMedian(t *testing.T) {
	for _, fn := range medianFunctions {
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

var (
	benchmarks = []int{100, 1000, 10000, 100000}
)

func BenchmarkMedian(b *testing.B) {
	for _, fn := range medianFunctions {
		b.Run(fn.name, func(b *testing.B) {
			for _, size := range benchmarks {
				b.Run(strconv.Itoa(size), func(b *testing.B) {
					// Static random numbers per run,
					// function and benchmark
					r := rand.New(rand.NewSource(int64(size)))
					list := make([]float64, size)
					for i := range list {
						list[i] = r.Float64()
					}
					for i := 0; i < b.N; i++ {
						_ = fn.fn(list)
					}
				})
			}

		})
	}
}
