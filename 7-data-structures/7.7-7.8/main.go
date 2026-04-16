package main

import "fmt"

func main() {
	userNames := make([]string, 2)
	userNames = append(userNames, "Tyler")
	userNames = append(userNames, "Jordan")

	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 2)
	courseRatings["Go"] = 4.7
	courseRatings["React"] = 4.5

	fmt.Println(courseRatings)

	for i, name := range userNames {
		fmt.Println(i, name)
	}

	for course, rating := range courseRatings {
		fmt.Println(course, rating)
	}
}
