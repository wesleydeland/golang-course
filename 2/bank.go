package main

import (
	"fmt"

	"example.com/bank/file_ops"
	"github.com/Pallinder/go-randomdata"
)

const fileName = "balance.txt"

func main() {
	var accountBalance, err = file_ops.ReadFloatFromFile(fileName)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		panic("FATAL ERROR")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 - ", randomdata.PhoneNumber())

	for {
		presentOptions()

		fmt.Print("Your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("How much are you depositing?:")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			file_ops.WriteFloatToFile(accountBalance, fileName)
			fmt.Println("Your balance is:", accountBalance)
		case 3:
			fmt.Print("How much are you withdrawing?:")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Withdrawal must be more than 0")
				continue
			}

			if withdrawAmount <= accountBalance {
				accountBalance -= withdrawAmount
				file_ops.WriteFloatToFile(accountBalance, fileName)
			} else {
				fmt.Println("Invalid amount.")
				continue
			}
			fmt.Println("Your balance is:", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing Go Bank")
			return
		}

	}

}
