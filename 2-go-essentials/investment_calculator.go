package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5

	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Enter an investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter a number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter an expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureValueInflationAdj := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValueInflationAdj)
}
