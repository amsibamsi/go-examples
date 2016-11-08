package radixsort

import "math"

func RadixSort(a []int) {
	var (
		counts [10]int
		buffer []int
		digit  []int
	)
	buffer = make([]int, len(a))
	digit = make([]int, len(a))
	if len(a) == 0 {
		return
	}
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	length := int(math.Log10(float64(max))) + 1
	for l := 0; l < length; l++ {
		for i := range counts {
			counts[i] = 0
		}
		for i, v := range a {
			n := v
			for i := 0; i < l; i++ {
				n /= 10
			}
			n %= 10
			digit[i] = n
			counts[n]++
		}
		for i := 1; i < len(counts); i++ {
			counts[i] += counts[i-1]
		}
		for i := len(a) - 1; i >= 0; i-- {
			counts[digit[i]] -= 1
			index := counts[digit[i]]
			buffer[index] = a[i]
		}
		for i, v := range buffer {
			a[i] = v
		}
	}
}
