package main

import "fmt"

// Pass by value
func process(x int) {
	x = 100
}

// Pass by reference (pointer)
func processPtr(x *int) {
	*x = 100
}

func main() {
	x := 1
	process(x)

	fmt.Println("x after process(x):", x) // x == 1 (unchanged)

	processPtr(&x)
	fmt.Println("x after processPtr(&x):", x) // x == 100 (changed)
}
