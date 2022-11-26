package main

import "fmt"

// constant of days in weeks
const (
	_ = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday = "Sunday"
)

func main() {
	fmt.Print(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
