package mergesort

func MergeSort(a []int) {
	length := len(a)
	buf := make([]int, length)
	src := a
	dst := buf
	for size := 1; size < length; size *= 2 {
		dstInd := 0
		for part := 0; part < length; part += 2 * size {
			left := part
			right := part + size
			for left < part+size && left < length && right < part+2*size && right < length {
				if src[left] <= src[right] {
					dst[dstInd] = src[left]
					left++
				} else {
					dst[dstInd] = src[right]
					right++
				}
				dstInd++
			}
			for left < part+size && left < length {
				dst[dstInd] = src[left]
				dstInd++
				left++
			}
			for right < part+2*size && right < length {
				dst[dstInd] = src[right]
				dstInd++
				right++
			}
		}
		src, dst = dst, src
	}
	copy(a, src)
}
