package main

import (
	"fmt"

	"com.analab/gobank/fileops"

	"github.com/Pallinder/go-randomdata"
)

const balanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromValue(balanceFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 27/7:", randomdata.PhoneNumber())

	for { //loop infinito
		ShowMenu()

		var operation int

		fmt.Scan(&operation)

		switch operation {
		case 1:
			fmt.Println("Your balance is: US$", accountBalance)
		case 2:
			var depositMoney float64
			fmt.Println("Deposit value:")
			fmt.Scan(&depositMoney)
			if depositMoney > 0 {
				accountBalance = accountBalance + depositMoney
				fmt.Println("Balance updated to US$:", accountBalance)
				fileops.WriteFloatToFile(accountBalance, balanceFile)
			} else {
				fmt.Println("Invalid ammount. Must be greater than 0")
				continue //skip
			}
		case 3:
			var withdraw float64
			fmt.Println("Withdraw value:")
			fmt.Scan(&withdraw)

			if withdraw < 0 {
				fmt.Println("Invalid ammount. Must be greater than 0")
				continue
			}

			if withdraw > accountBalance {
				fmt.Println("Not enough balance.")
				continue
			}

			accountBalance = accountBalance - withdraw
			fmt.Println("Balance updated to US$:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, balanceFile)
		case 4:
			fmt.Println("Thanks for choosing Go Bank!")
			return
		default:
			fmt.Println("Invalid option")
		}
	}

}
