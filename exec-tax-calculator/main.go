package main

import (
	"analab/golang/prices"
)

func main() {
	taxRates := []float64{0.05, 0.10, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}
