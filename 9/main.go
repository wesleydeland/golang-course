package main

import (
	"fmt"

	filemanager "example.com/price-calculator/fileManager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0.8, 0.07, 0.010, 0.50}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result%.0f.json", taxRate*100))
		//cmd := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(taxRate, fm)
		err := job.Process()

		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
