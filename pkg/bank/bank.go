package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceTxt), 0644)

}

func main() {
	var accountBalance = getAccountBalance()

	for {
		fmt.Println("Welcome to go bank")
		fmt.Println("What you want to do")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Println("Enter your choice :")
		fmt.Scan(&choice)

		if choice > 4 {
			fmt.Println("your choice is invalid")
		}

		switch choice {
		case 1:
			balance := getAccountBalance()
			fmt.Println("Your account balance is ", balance)
		case 2:
			var depositedAmount float64
			fmt.Println("Enter your deposit amount ")
			fmt.Scan(&depositedAmount)

			if depositedAmount <= 0 {
				fmt.Println("Invalid amount, Must be greater than 0")
				continue
			}
			
			accountBalance += depositedAmount
			fmt.Println("Your updated account balance : ", accountBalance)
			writeBalanceToFile(accountBalance)

		case 3:
			var withdrawalAmount float64
			fmt.Println("Enter withdrawal amount ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount, Must be greater than 0")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("You can't withdraw more then your account balance")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Your updated account balance : ", accountBalance)
			writeBalanceToFile(accountBalance)

		default:
			fmt.Println("Good Bye..!")
			fmt.Println("Thanks for choosing go bank...!")
			return
		}
	}

}

func getAccountBalance() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceString := string(data)
	balanceValue, _ :=strconv.ParseFloat(balanceString, 64)
	return balanceValue
}
