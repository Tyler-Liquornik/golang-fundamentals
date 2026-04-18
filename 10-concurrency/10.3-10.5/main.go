package main

import (
	"fmt"

	"github.com/Tyler-Liquornik/golang-fundamentals/10-concurrency/10.3-10.5/filemanager"
	"github.com/Tyler-Liquornik/golang-fundamentals/10-concurrency/10.3-10.5/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)

		inputPath := "10-concurrency/10.3-10.5/prices/prices.txt"
		outputPath := fmt.Sprintf("10-concurrency/10.3-10.5/results/result_%.0f.json", taxRate*100)
		fm := filemanager.New(inputPath, outputPath)

		// cmdm := cmdmanager.New() can get input from cmdm now instead of from files

		priceJob := prices.NewTaxIncludedPriceJob(fm /*cmdm*/, taxRate)
		go priceJob.Process(doneChans[i], errorChans[i])
	}

	for i := range taxRates {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Println("Error processing prices:", err)
			}
		case <-doneChans[i]:
			fmt.Println("Prices processed successfully!")
		}
	}
}
