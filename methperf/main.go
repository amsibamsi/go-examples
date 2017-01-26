package methperf

type Slice struct {
	slice []int
}

func NewSlice(size, cap int) Slice {
	s := make([]int, size, cap)
	return Slice{s}
}

func (s *Slice) Append(v int) {
	s.slice = append(s.slice, v)
}

func (s *Slice) Set(i, v int) {
	s.slice[i] = v
}

func (s *Slice) Get(i int) int {
	return s.slice[i]
}
