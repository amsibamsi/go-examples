package cgobench

// #include <cgobench.h>
import "C"

func GoFunction(n int) int {
	return n + 1
}

func CallGo(n int) int {
	return GoFunction(n)
}

func CallCGO(n int) int {
	return int(C.c_function(C.int(n)))
}
