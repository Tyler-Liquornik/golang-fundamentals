package main

import "fmt"

type transformFunction func(*int)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	transformNumbers(nums, getTransformerFunction("double"))
	transformNumbers(nums, getTransformerFunction("subtractOne"))
	fmt.Println(nums)
}

func transformNumbers(nums []int, transform transformFunction) {
	for i := range nums {
		transform(&nums[i])
	}
}

func getTransformerFunction(name string) transformFunction {
	switch name {
	case "double":
		return double
	case "subtractOne":
		return subtractOne
	default:
		return nil
	}
}

func double(number *int) {
	*number *= 2
}

func subtractOne(number *int) {
	*number -= 1
}
