package main

import (
	"fmt"
	"os"
)

const fileName = "output.txt"

//goals
//validate user input -> show error message and exit if invalid input is provided
// no negative numbers
// not 0

//store calculated restults into file

func main() {
	var revenue, expenses, taxRate float64
	var inputError1, inputError2, inputError3 error
	// fmt.Print("Please enter revenue: ")
	// fmt.Scan(&revenue)
	inputError1 = promptAndGetInput("Please enter revenue: ", &revenue)
	inputError2 = promptAndGetInput("Please enter expenses: ", &expenses)
	inputError3 = promptAndGetInput("Please enter tax rate: ", &taxRate)

	if inputError1 != nil || inputError2 != nil || inputError3 != nil {
		fmt.Println("invalid input.")
		return
	} else if revenue <= 0 || expenses <= 0 || taxRate <= 0 {
		fmt.Println("must be greater than 0")
		return
	}

	// fmt.Print("Please enter expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("Please enter tax rate: ")
	// fmt.Scan(&taxRate)

	// ebt := revenue - expenses
	// profit := ebt - (ebt * (taxRate / 100))
	// ratio := ebt / profit
	ebt, profit, ratio := calculateValues(revenue, expenses, taxRate)

	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
	writeResultsToFile(ebt, profit, ratio)
}

func promptAndGetInput(prompt string, input *float64) error {
	fmt.Print(prompt)
	_, err := fmt.Scan(input)
	return err
}

func calculateValues(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - (ebt * (taxRate / 100))
	ratio = ebt / profit
	return ebt, profit, ratio
}

func writeResultsToFile(ebt, profit, ratio float64) {
	resultText := fmt.Sprintf("EBT: %.2f\n", ebt)
	resultText += fmt.Sprintf("Profit: %.2f\n", profit)
	resultText += fmt.Sprintf("Ratio: %.2f\n", ratio)

	os.WriteFile(fileName, []byte(resultText), 0644)
}
