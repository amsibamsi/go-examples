package methperf

import "testing"

var size = 1000

func TestSliceStruct(t *testing.T) {
	s := []int{1, 2, 3}
	r := Slice{}
	for _, v := range s {
		r.Append(v)
	}
	for i := range s {
		v := r.Get(i)
		if v != s[i] {
			t.Errorf("Got value '%v' at index '%v', expected value '%v'", v, i, s[i])
		}
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, size, size)
		for j := 0; j < size; j++ {
			s[j] = j
		}
	}
}

func BenchmarkSliceStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := NewSlice(size, size)
		for j := 0; j < size; j++ {
			s.Set(j, j)
		}
	}
}
