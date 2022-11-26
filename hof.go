package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func minus(x, y int) int {
	return x - y
}

type Func func(int, int) int

func compute(xs Func) func(y int) int {
	x := xs(5, 2)
	f := func(y int) int {
		return x * y
	}

	return f
}

func main() {
	ms := compute(minus)
	fmt.Println("R:", ms(2))

	ad := compute(add)
	fmt.Println("R2:", ad(2))
}
