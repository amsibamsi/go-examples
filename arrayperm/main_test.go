package arrayperm

import (
	"fmt"
	"testing"
)

func TestArrayPermutations(t *testing.T) {
	tests := []struct {
		list  [][]int
		perms [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4},
				{5, 6},
			},
			[][]int{
				{1, 4, 5},
				{1, 4, 6},
				{2, 4, 5},
				{2, 4, 6},
				{3, 4, 5},
				{3, 4, 6},
			},
		},
		{
			[][]int{},
			[][]int{},
		},
		{
			[][]int{{}},
			[][]int{},
		},
		{
			[][]int{
				{},
				{1, 2},
			},
			[][]int{
				{1},
				{2},
			},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.list), func(t *testing.T) {
			p := ArrayPermutations(test.list)
			// Will test for exact ordering of permutations
			if len(p) != len(test.perms) {
				t.Errorf("got '%v' permutations, expected '%v'", len(p), len(test.perms))
			}
			for i := range p {
				if len(p[i]) != len(test.perms[i]) {
					t.Errorf("permutation '%v' has length '%v', expected '%v'", i, len(p[i]), len(test.perms[i]))
				}
				for j := range p[i] {
					if p[i][j] != test.perms[i][j] {
						t.Errorf("permutation '%v' is '%v', expected '%v'", i, p[i], test.perms[i])
					}
				}
			}
		})
	}
}
