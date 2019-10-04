package median

import (
	"sort"
)

func medianSort(list []float64) float64 {
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

func medianKSelect(l []float64) float64 {
	if len(l) == 0 {
		return 0
	}
	list := l
	k := len(list) / 2
	if len(list)%2 == 1 {
		k++
	}
	pivotIndex := 0
	for {
		left := make([]float64, 0)
		pivots := make([]float64, 0)
		right := make([]float64, 0)
		// Rotate pivot to not get stuck
		pivotIndex = (pivotIndex + 1) % len(list)
		pivot := list[pivotIndex]
		for _, n := range list {
			if n < pivot {
				left = append(left, n)
				continue
			}
			if n == pivot {
				pivots = append(pivots, n)
				continue
			}
			right = append(right, n)
		}
		//fmt.Printf("%v\nk=%v\npivot=%v\n%v|%v|%v\n\n", list, k, pivot, left, pivots, right)
		if k <= len(left) {
			list = left
		}
		if k > len(left) && k <= len(left)+len(pivots) {
			return pivot
		}
		if k > len(left)+len(pivots) {
			list = right
			k = k - len(left) - len(pivots)
		}
	}
}
