package main

import "fmt"

func main() {
	result := add(1, 2)
	resultPlus1 := result + 1
	fmt.Println(resultPlus1)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
