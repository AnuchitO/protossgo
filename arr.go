package main

import "fmt"

func hello(ts []string) {
	fmt.Printf("%#v\n", ts)
}

func main() {
	var xs []string
	xs = append(xs, "AnuchitO")
	_ = append(xs, "nong", "nong", "nong")
	// xs[0] = "Anuchit"
	// fmt.Println(ys)
	l := len(xs)

	hello(xs)

	fmt.Println("len:", l, cap(xs))

	fmt.Printf("%#v\n", xs)
	// fmt.Printf("%v\n", xs[0])

}
