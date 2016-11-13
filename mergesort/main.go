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
