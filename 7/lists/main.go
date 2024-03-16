package main

import "fmt"

// type Product struct {
// 	title string
// 	id    string
// 	price float64
// }

// type temperatureData struct {
// 	day1 float64
// }

// func main() {
// 	var productNames [4]string
// 	productNames[3] = "cookies"
// 	prices := [4]float64{1.15, 2.25, 3.35, 4.45}
// 	fmt.Println(prices[2])
// 	fmt.Println(productNames)

// 	featuredPrices := prices[:3]
// 	highlightedPrices := featuredPrices[1:]
// 	fmt.Println(featuredPrices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// 	fmt.Println(highlightedPrices)
// }

func main() {
	prices := []float64{10.99, 11.99}

	prices = append(prices, 12.99, 13.99)

	fmt.Println(prices)

	discountPrices := []float64{1.22, 2.33}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}
