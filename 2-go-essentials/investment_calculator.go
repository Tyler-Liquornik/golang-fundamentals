package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount, years, expectedReturnRate float64

	outputText("Enter an investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter a number of years: ")
	fmt.Scan(&years)

	outputText("Enter an expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureValueInflationAdj := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for inflation): %.2f", futureValue, futureValueInflationAdj)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, fvAdj float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fvAdj = fv / math.Pow(1+inflationRate/100, years)
	return
}
