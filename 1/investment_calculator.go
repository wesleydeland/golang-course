package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	outputText("Please enter an investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Please enter the expected rate of return: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Please enter the number of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	//formattedFRV := fmt.Sprintf("Future Value (Adjusted for Inflation): %.2f\n", futureRealValue)

	//outputs information
	//fmt.Println("Future Value: ", futureValue)
	//fmt.Print(formattedFV, formattedFRV)
	//fmt.Println("Future Value (Adjusted for Inflation): ", futureRealValue)
	fmt.Printf(
		`Future Value: %.2f
Future Value (Adjusted for Inflation): %.2f`, futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
