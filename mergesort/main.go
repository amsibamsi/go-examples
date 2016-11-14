package mergesort

import (
	"fmt"
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

func merge(a, buf []int, start, end int, recur bool) {
	length := end - start + 1
	src := a
	dst := buf
	startSize := 1
	if !recur {
		for startSize < length/2 {
			startSize *= 2
		}
	}
	fmt.Printf("merge %v-%v, length %v, size %v\n", start, end, length, startSize)
	for size := startSize; size < length; size *= 2 {
		fmt.Printf("merge size loop %v\n", size)
		dstInd := start
		for part := start; part < end; part += 2 * size {
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
	fmt.Printf("merged %v-%v, res %v\n", start, end, src[start:end+1])
}

func recMergeSort(a, buf []int, start, end, tasks int) {
	fmt.Printf("mergesort %v-%v\n", start, end)
	if (end - start) > 0 {
		if tasks < runtime.NumCPU() {
			wg := sync.WaitGroup{}
			mid := start + (end-start)/2
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
			merge(a, buf, start, end, false)
		} else {
			merge(a, buf, start, end, true)
		}
	}
}

func RecMergeSort(a []int) {
	buf := make([]int, len(a))
	recMergeSort(a, buf, 0, len(a)-1, 1)
}
