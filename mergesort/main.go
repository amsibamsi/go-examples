package mergesort

import (
	"runtime"
	"sync"
)

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

func ParMergeSort(a []int) {
	length := len(a)
	buf := make([]int, length)
	src := a
	dst := buf
	wg := sync.WaitGroup{}
	for size := 1; size < length; size *= 2 {
		for part := 0; part < length; part += 2 * size {
			wg.Add(1)
			go func(part int) {
				dstInd := part
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
				wg.Done()
			}(part)
		}
		wg.Wait()
		src, dst = dst, src
	}
	copy(a, src)
}

func Par2MergeSort(a []int) {
	cpus := runtime.NumCPU()
	length := len(a)
	buf := make([]int, length)
	src := a
	dst := buf
	workers := sync.WaitGroup{}
	jobs := make(chan struct{ part, size int }, cpus)
	results := make(chan struct{}, cpus)
	done := sync.WaitGroup{}
	for i := 0; i < cpus; i++ {
		workers.Add(1)
		go func() {
			for job := range jobs {
				part := job.part
				size := job.size
				dstInd := part
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
				results <- struct{}{}
			}
			workers.Done()
		}()
	}
	for size := 1; size < length; size *= 2 {
		done.Add(1)
		go func() {
			for p := 0; p < length; p += 2 * size {
				<-results
			}
			done.Done()
		}()
		for part := 0; part < length; part += 2 * size {
			jobs <- struct{ part, size int }{part, size}
		}
		done.Wait()
		src, dst = dst, src
	}
	close(jobs)
	workers.Wait()
	close(results)
	copy(a, src)
}

func recMerge(a, buf []int, start, end int) {
	length := end - start + 1
	src := a
	dst := buf
	for size := 1; size < length; size *= 2 {
		dstInd := start
		for part := start; part <= end; part += 2 * size {
			left := part
			right := part + size
			for left < part+size && left <= end && right < part+2*size && right <= end {
				if src[left] <= src[right] {
					dst[dstInd] = src[left]
					left++
				} else {
					dst[dstInd] = src[right]
					right++
				}
				dstInd++
			}
			for left < part+size && left <= end {
				dst[dstInd] = src[left]
				dstInd++
				left++
			}
			for right < part+2*size && right <= end {
				dst[dstInd] = src[right]
				dstInd++
				right++
			}
		}
		src, dst = dst, src
	}
	copy(a[start:end+1], src[start:end+1])
}

func merge(a, buf []int, start, mid, end int) {
	left := start
	right := mid + 1
	ind := start
	for left <= mid && right <= end {
		if a[left] <= a[right] {
			buf[ind] = a[left]
			left++
		} else {
			buf[ind] = a[right]
			right++
		}
		ind++
	}
	for left <= mid {
		buf[ind] = a[left]
		ind++
		left++
	}
	for right <= end {
		buf[ind] = a[right]
		ind++
		right++
	}
	copy(a[start:end+1], buf[start:end+1])
}

func recMergeSort(a, buf []int, start, end, tasks int) {
	if (end - start) > 0 {
		if tasks < runtime.NumCPU() {
			mid := start + (end-start)/2
			wg := sync.WaitGroup{}
			wg.Add(2)
			go func() {
				recMergeSort(a, buf, start, mid, tasks*2)
				wg.Done()
			}()
			go func() {
				recMergeSort(a, buf, mid+1, end, tasks*2)
				wg.Done()
			}()
			wg.Wait()
			merge(a, buf, start, mid, end)
		} else {
			recMerge(a, buf, start, end)
		}
	}
}

func RecMergeSort(a []int) {
	buf := make([]int, len(a))
	recMergeSort(a, buf, 0, len(a)-1, 1)
}
