package pairsum

import (
	"fmt"
	"testing"
)

var (
	Tests = []struct {
		list   []int
		sum    int
		result bool
	}{
		{[]int{1, 2, 3, 9}, 8, false},
		{[]int{1, 2, 4, 4}, 8, true},
	}
)

func TestHasPairWithSum(t *testing.T) {
	for _, test := range Tests {
		t.Run(fmt.Sprintf("%v", test.list), func(t *testing.T) {
			r := HasPairWithSum(test.list, test.sum)
			if r != test.result {
				t.Errorf("got '%v', expected '%v'", r, test.result)
			}
		})
	}
}
