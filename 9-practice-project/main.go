package main

import (
	"fmt"

	"github.com/Tyler-Liquornik/golang-fundamentals/9-practice-project/filemanager"
	"github.com/Tyler-Liquornik/golang-fundamentals/9-practice-project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		inputPath := "9-practice-project/prices/prices.txt"
		outputPath := fmt.Sprintf("9-practice-project/results/result_%.0f.json", taxRate*100)
		fm := filemanager.New(inputPath, outputPath)

		// cmdm := cmdmanager.New() can get input from cmdm now instead of from files

		priceJob := prices.NewTaxIncludedPriceJob(fm /*cmdm*/, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Error processing prices:", err)
			continue
		}
	}
}
