package median

import (
	"fmt"
	"sort"
)

func median(list []float64) float64 {
	l := len(list)
	if l == 0 {
		return 0
	}
	sorted := make([]float64, l)
	copy(sorted, list)
	sort.Float64s(sorted)
	if l%2 == 1 {
		return sorted[l/2]
	}
	return (sorted[l/2] + sorted[l/2-1]) / 2
}

func median2(list []float64) float64 {
	if len(list) == 0 {
		return 0
	}
	l := list
	k := len(list) / 2
	pi := 0
	for i := 10; i > 0; i-- {
		left := make([]float64, 0)
		right := make([]float64, 0)
		p := l[pi%len(l)]
		pi++
		for _, n := range l {
			if n <= p {
				left = append(left, n)
				continue
			}
			right = append(right, n)
		}
		fmt.Printf("i=%v\nl: %v\nk=%v\np=%v\nlr: %v|%v\n\n", i, l, k, p, left, right)
		if len(left) == len(right) {
			return p
		}
		if len(left) == len(right)+1 {
			return p
		}
		if len(left) > len(right)+1 {
			l = left
			continue
		}
		//if len(right) > k
		l = right
		k = k - len(left)
	}
	return 0
}
