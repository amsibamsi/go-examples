package revstr

import "testing"

var (
	Tests = []struct {
		original, reversed string
	}{
		{"", ""},
		{"x", "x"},
		{"ab", "ba"},
		{"my string", "gnirts ym"},
		{"ê獅", "獅ê"},
	}
	Benchmark = "abcdefghijklmnopqrtsuvwxyz1234567890"
)

type Reverser func(string) string

func test(r Reverser, t *testing.T) {
	for _, test := range Tests {
		t.Run(test.original, func(t *testing.T) {
			s := r(test.original)
			if s != test.reversed {
				t.Errorf("got '%s', expected '%s'", s, test.reversed)
			}
		})
	}
}

func TestReverseStringBuffer(t *testing.T) {
	test(ReverseStringBuffer, t)
}

func TestReverseStringAppend(t *testing.T) {
	test(ReverseStringAppend, t)
}

func TestReverseStringIndex(t *testing.T) {
	test(ReverseStringIndex, t)
}

func TestReverseStringInplace(t *testing.T) {
	test(ReverseStringInplace, t)
}

func benchmark(r Reverser, b *testing.B) {
	for i := 0; i < b.N; i++ {
		r(Benchmark)
	}
}

func BenchmarkReverseStringBuffer(b *testing.B) {
	benchmark(ReverseStringBuffer, b)
}

func BenchmarkReverseStringAppend(b *testing.B) {
	benchmark(ReverseStringAppend, b)
}

func BenchmarkReverseStringIndex(b *testing.B) {
	benchmark(ReverseStringIndex, b)
}

func BenchmarkReverseStringInplace(b *testing.B) {
	benchmark(ReverseStringInplace, b)
}
