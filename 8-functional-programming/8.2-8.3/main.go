package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	doubleTransformer := createTransformer(2)
	transformed := transformNumbers(&numbers, doubleTransformer)

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}
