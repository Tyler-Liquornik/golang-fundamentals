package main

import (
	"fmt"

	"github.com/Tyler-Liquornik/golang-fundamentals/3-packages/file_io"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = file_io.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	for {
		presentOptions()

		fmt.Print("Your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your new balance is", accountBalance)
			file_io.WriteFloatToFile(accountBalanceFile, accountBalance)
		case 3:
			fmt.Print("Your withdrawal: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount > accountBalance {
				fmt.Println("You don't have enough money!")
				continue
			} else if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Your new balance is", accountBalance)
			file_io.WriteFloatToFile(accountBalanceFile, accountBalance)
		default:
			fmt.Println("Goodbye!")

			fmt.Println("Your choice:", choice)
			fmt.Println("Thank you for using Go Bank!")
			return
		}
	}
}
