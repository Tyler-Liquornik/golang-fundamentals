package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	updatedPrices := append(prices, 12.99)

	fmt.Println(prices)
	fmt.Println(updatedPrices)

	arr := [5]float64{10.99, 8.99}
	pricesWithExtraCap := arr[:2]
	pricesWithExtraCap = append(pricesWithExtraCap, 12.99)

	fmt.Println(arr)
	fmt.Println(pricesWithExtraCap)
}
