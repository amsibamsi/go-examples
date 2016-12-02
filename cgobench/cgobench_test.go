package cgobench

import "testing"

func BenchmarkGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := CallGo(i)
		if n != i+1 {
			b.Fail()
		}
	}
}

func BenchmarkCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := CallCGO(i)
		if n != i+1 {
			b.Fail()
		}
	}
}
