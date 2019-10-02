package median

import "sort"

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

// TODO: review
func median2(list []float64) float64 {
	if len(list) == 0 {
		return 0
	}
	l := list
	k := len(list) / 2
	for {
		left := make([]float64, len(l))
		right := make([]float64, len(l))
		p := l[0]
		for _, n := range l {
			if n <= p {
				left = append(left, n)
			} else {
				right = append(right, n)
			}
		}
		if len(left)-1 == k {
			return p
		}
		if len(left) == len(right) {
			return p
		}
		if len(left) > k+1 {
			l = left
			continue
		}
		//if len(right) > k
		l = right
		k = k - len(left)
	}
}
