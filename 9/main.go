package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.010, 0.50}

	result := make(map[float64][]float64, 12)

	for _, taxRate := range taxRates {
		result[taxRate] = calculateTax(taxRate, prices)
	}

	fmt.Println(result)
}

func calculateTax(taxRate float64, prices []float64) []float64 {
	results := []float64{}

	for _, price := range prices {
		results = append(results, price*(1+taxRate))
	}

	return results
}
