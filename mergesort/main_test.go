package mergesort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

var (
	MaxInt = int(^uint(0) >> 1)
	Tests  = []struct {
		original, sorted []int
	}{
		{[]int{3, 8, 7, 5, 6, 1, 9, 2, 4}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 1, 1, 2, 1e10, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1, 2, 1e10}},
	}
	BenchmarkLength = int(1e5)
)

func randArray(length, max int) []int {
	r := make([]int, length)
	for i := 0; i < len(r); i++ {
		r[i] = rand.Intn(max)
	}
	return r
}

func TestMergeSort(t *testing.T) {
	for _, test := range Tests {
		t.Run(fmt.Sprintf("%v", test.original), func(t *testing.T) {
			l := make([]int, len(test.original))
			copy(l, test.original)
			MergeSort(l)
			for i := range test.sorted {
				if l[i] != test.sorted[i] {
					t.Fatalf("got '%v', expected '%v'", l, test.sorted)
				}
			}
		})
	}
}

func TestMergeSortRandom(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 1e3; i++ {
		t.Run(fmt.Sprintf("random %v", i), func(t *testing.T) {
			a := randArray(1e3, MaxInt)
			b := make([]int, len(a))
			copy(b, a)
			sort.Ints(a)
			MergeSort(b)
			for j := range a {
				if b[j] != a[j] {
					t.Fatalf("got '%v', expected '%v'", b[j], a[j])
				}
			}
		})
	}
}

func TestParMergeSort(t *testing.T) {
	for _, test := range Tests {
		t.Run(fmt.Sprintf("%v", test.original), func(t *testing.T) {
			l := make([]int, len(test.original))
			copy(l, test.original)
			ParMergeSort(l)
			for i := range test.sorted {
				if l[i] != test.sorted[i] {
					t.Fatalf("got '%v', expected '%v'", l, test.sorted)
				}
			}
		})
	}
}

func TestParMergeSortRandom(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 1e3; i++ {
		t.Run(fmt.Sprintf("random %v", i), func(t *testing.T) {
			a := randArray(1e3, MaxInt)
			b := make([]int, len(a))
			copy(b, a)
			sort.Ints(a)
			ParMergeSort(b)
			for j := range a {
				if b[j] != a[j] {
					t.Fatalf("got '%v', expected '%v'", b[j], a[j])
				}
			}
		})
	}
}

func TestPar2MergeSort(t *testing.T) {
	for _, test := range Tests {
		t.Run(fmt.Sprintf("%v", test.original), func(t *testing.T) {
			l := make([]int, len(test.original))
			copy(l, test.original)
			Par2MergeSort(l)
			for i := range test.sorted {
				if l[i] != test.sorted[i] {
					t.Fatalf("got '%v', expected '%v'", l, test.sorted)
				}
			}
		})
	}
}

func TestPar2MergeSortRandom(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 1e3; i++ {
		t.Run(fmt.Sprintf("random %v", i), func(t *testing.T) {
			a := randArray(1e3, MaxInt)
			b := make([]int, len(a))
			copy(b, a)
			sort.Ints(a)
			Par2MergeSort(b)
			for j := range a {
				if b[j] != a[j] {
					t.Fatalf("got '%v', expected '%v'", b[j], a[j])
				}
			}
		})
	}
}

func TestRecMergeSort(t *testing.T) {
	for _, test := range Tests {
		t.Run(fmt.Sprintf("%v", test.original), func(t *testing.T) {
			l := make([]int, len(test.original))
			copy(l, test.original)
			RecMergeSort(l)
			for i := range test.sorted {
				if l[i] != test.sorted[i] {
					t.Fatalf("got '%v', expected '%v'", l, test.sorted)
				}
			}
		})
	}
}

func TestRecMergeSortRandom(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 1e3; i++ {
		t.Run(fmt.Sprintf("random %v", i), func(t *testing.T) {
			a := randArray(1e3, MaxInt)
			b := make([]int, len(a))
			copy(b, a)
			sort.Ints(a)
			RecMergeSort(b)
			for j := range a {
				if b[j] != a[j] {
					t.Fatalf("got '%v', expected '%v'", b[j], a[j])
				}
			}
		})
	}
}

func BenchmarkSort(b *testing.B) {
	rand.Seed(0)
	for i := 0; i < b.N; i++ {
		a := randArray(BenchmarkLength, MaxInt)
		sort.Ints(a)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	rand.Seed(0)
	for i := 0; i < b.N; i++ {
		a := randArray(BenchmarkLength, MaxInt)
		MergeSort(a)
	}
}

func BenchmarkParMergeSort(b *testing.B) {
	rand.Seed(0)
	for i := 0; i < b.N; i++ {
		a := randArray(BenchmarkLength, MaxInt)
		ParMergeSort(a)
	}
}

func BenchmarkPar2MergeSort(b *testing.B) {
	rand.Seed(0)
	for i := 0; i < b.N; i++ {
		a := randArray(BenchmarkLength, MaxInt)
		Par2MergeSort(a)
	}
}

func BenchmarkRecMergeSort(b *testing.B) {
	rand.Seed(0)
	for i := 0; i < b.N; i++ {
		a := randArray(BenchmarkLength, MaxInt)
		RecMergeSort(a)
	}
}
