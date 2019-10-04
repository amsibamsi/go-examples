package main

import "fmt"

type pardoFn func(int, int) int

type doubler interface {
	double(int) int
}

func computeDouble(fn doubler, x int) int {
	return fn.double(x)
}

func mul(x, y int) int {
	return x * y
}

func (f pardoFn) double(x int) int {
	return f(x, 2)
}

func main() {
	fmt.Println(computeDouble(pardoFn(mul), 2))
}
