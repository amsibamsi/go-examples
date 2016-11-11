package mergesort

import (
	"fmt"
	"testing"
)

var (
	Tests = []struct {
		original, sorted []int
	}{
		{[]int{3, 8, 7, 5, 6, 1, 9, 2, 4}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 1, 1, 2, 1e10, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1, 2, 1e10}},
	}
)

func TestMergeSort(t *testing.T) {
	for _, test := range Tests {
		t.Run(fmt.Sprintf("%v", test.original), func(t *testing.T) {
			l := make([]int, len(test.original))
			copy(l, test.original)
			MergeSort(l)
			for i := range test.sorted {
				if l[i] != test.sorted[i] {
					t.Fatalf("got '%v', expected '%v'", l, test.sorted)
				}
			}
		})
	}
}
