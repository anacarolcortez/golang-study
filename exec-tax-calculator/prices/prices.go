package prices

import (
	"fmt"
	"strconv"

	"analab/golang/fileManager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	Prices            []float64
	TaxIncludedPrices map[string]string
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	prices := getPricesFromFile("prices.txt")
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
		Prices:  prices,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.calculate()
	job.saveTaxedPricesToJson()
}

func (job *TaxIncludedPriceJob) calculate() {
	job.TaxIncludedPrices = make(map[string]string, len(job.Prices))
	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		key := fmt.Sprintf("%.2f", price)
		job.TaxIncludedPrices[key] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
}

func getPricesFromFile(fileName string) []float64 {
	line, err := fileManager.ReadTxtFile(fileName)

	if err != nil {
		fmt.Println("Error getting prices from file")
		return []float64{}
	}

	var prices []float64
	for index, priceString := range line {
		floatPrice, err := strconv.ParseFloat(priceString, 64)
		if err != nil {
			fmt.Printf("Error converting price to float in line %d\n", index+1)
			return []float64{}
		}
		prices = append(prices, floatPrice)
	}

	return prices
}

func (job *TaxIncludedPriceJob) saveTaxedPricesToJson() {
	path := fmt.Sprintf("taxedPrices_%.0f.json", job.TaxRate*100)
	err := fileManager.WriteJSONFile(path, job)
	if err != nil {
		fmt.Println("Error saving taxed prices to file")
	}
}
