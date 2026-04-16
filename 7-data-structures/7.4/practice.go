package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	hobbies := [3]string{"Coding", "Traveling", "Partying"}
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	slice1 := hobbies[0:2]
	slice2 := hobbies[:2]
	fmt.Println(slice1, slice2)

	slice1 = hobbies[1:]
	fmt.Println(slice1)

	goals := []string{"Learn Go", "Work with Go"}
	goals[1] = "Work with Go at Scale"
	goals = append(goals, "Build AI Systems in Go")
	fmt.Println(goals)

	product := Product{
		title: "Go Gopher",
		id:    "1234",
		price: 19.99,
	}
	product2 := Product{
		title: "Duke",
		id:    "5678",
		price: 99.99,
	}

	products := []Product{product, product2}
	products = append(products, Product{"Tux", "9012", 999.99})

	fmt.Println(products)
}
