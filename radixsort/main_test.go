package radixsort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

var (
	MaxInt = int(^uint(0) >> 1)
)

func array(length, max int, seed int64) []int {
	rand.Seed(seed)
	r := make([]int, length)
	for i := 0; i < len(r); i++ {
		r[i] = rand.Intn(max)
	}
	return r
}

func randomArray(length, max int) []int {
	return array(length, max, time.Now().UTC().UnixNano())
}

func testMany(length, max, cases int, random bool, t *testing.T) {
	for i := 0; i < cases; i++ {
		var a []int
		if random {
			a = randomArray(length, max)
		} else {
			a = array(length, max, int64(i))
		}
		t.Run(fmt.Sprintf("array %v", i), func(t *testing.T) {
			r := make([]int, len(a))
			n := copy(r, a)
			if n != len(a) {
				t.Fatalf("could not copy whole array")
			}
			sort.Ints(a)
			RadixSort(r)
			for i := range a {
				if r[i] != a[i] {
					t.Fatalf("got '%v', expected '%v'", r[i], a[i])
				}
			}
		})
	}
}

func TestSimple(t *testing.T) {
	tests := []struct {
		original, sorted []int
	}{
		{[]int{}, []int{}},
		{[]int{4}, []int{4}},
		{[]int{4, 6, 5, 3, 7, 8, 1, 2, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{2e12, 1, 9e18}, []int{1, 2e12, 9e18}},
		{[]int{2, 7, 0}, []int{0, 2, 7}},
		{[]int{0, 1, 0, 1, 1, 1, 0}, []int{0, 0, 0, 1, 1, 1, 1}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.original), func(t *testing.T) {
			o := test.original
			RadixSort(o)
			for i := range o {
				if o[i] != test.sorted[i] {
					t.Fatalf("got '%v', expected '%v'", o, test.sorted)
				}
			}
		})
	}
}

func TestSmall(t *testing.T) {
	testMany(100, 1000, 1000, false, t)
}

func TestBig(t *testing.T) {
	testMany(1000, MaxInt, 1000, false, t)
}

func TestRandom(t *testing.T) {
	testMany(1000, MaxInt, 1000, true, t)
}

func benchmark(length, max int, customSort bool, b *testing.B) {
	a := array(length, max, 1)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c := make([]int, len(a))
		n := copy(c, a)
		if n != len(a) {
			b.Fatalf("could not copy whole array")
		}
		if customSort {
			RadixSort(c)
		} else {
			sort.Ints(c)
		}
	}
}

func BenchmarkFewSmall(b *testing.B) {
	benchmark(100, 1000, true, b)
}

func BenchmarkFewSmallCompare(b *testing.B) {
	benchmark(100, 1000, false, b)
}

func BenchmarkFewBig(b *testing.B) {
	benchmark(100, MaxInt, true, b)
}

func BenchmarkFewBigCompare(b *testing.B) {
	benchmark(100, MaxInt, false, b)
}

func BenchmarkManySmall(b *testing.B) {
	benchmark(10000, 1000, true, b)
}

func BenchmarkManySmallCompare(b *testing.B) {
	benchmark(10000, 1000, false, b)
}

func BenchmarkManyBig(b *testing.B) {
	benchmark(10000, MaxInt, true, b)
}

func BenchmarkManyBigCompare(b *testing.B) {
	benchmark(10000, MaxInt, false, b)
}
