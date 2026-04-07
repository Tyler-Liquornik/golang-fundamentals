package main

import "fmt"

func main() {
	x := 1
	y := x // copy

	y = 100

	fmt.Println("x:", x) // x == 1 (unchanged)
	fmt.Println("y:", y) // y == 100

	p := &x
	*p = 100

	fmt.Println("x after *p = 100:", x) // x == 100 (changed)
}
