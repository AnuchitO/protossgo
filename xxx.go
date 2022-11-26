package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	nong := "nong"
	const pi = 3.14

	fmt.Println(i, j, c, python, java, nong)

	nong = "AnuchitO"
	fmt.Println(i, j, c, python, java, nong)
}
