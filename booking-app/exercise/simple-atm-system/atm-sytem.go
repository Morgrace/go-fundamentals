package main

import (
	"fmt"
	"os"
)

var userBalance int = 1000
var userSelectedOption int
var amountDeposited int
var amountToWithraw int

// VALIDATION
func validateUserInput(userInput int) bool {
	if userInput >= 1 && userInput <= 4 {
		return true
	}
	fmt.Printf("🛑 %v is an invalid option. Please select a valid option\n\n", userInput)
	return false
}

// BANKING OPERATIONS
func showBalance() {
	fmt.Printf("Your balance: $%v\n\n", userBalance)
}

func depositMoney() {
	fmt.Print("Enter amount($): ")
	fmt.Scan(&amountDeposited)

	if amountDeposited <= 0 {
		fmt.Printf("💰$%v is an invalid amount. Enter a valid amount\n\n", amountDeposited)
		return
	}
	userBalance += amountDeposited

	fmt.Printf("💰$%v was deposited successfully!🥳 New balance: $%v\n\n", amountDeposited, userBalance)
}

func withdrawMoney() {
	fmt.Print("Enter amount($): ")
	fmt.Scan(&amountToWithraw)
	// Validate user input
	if amountToWithraw <= 0 {
		fmt.Println("\n🛑 Invalid amount")
		return
	}
	if amountToWithraw > userBalance {
		fmt.Println("🛑 Insufficient funds")
		return
	}
	userBalance -= amountToWithraw

	fmt.Printf("Withdrawal succesful!🥳 New balance: $%v\n\n", userBalance)
}

// ATM SYSTEM
func showMenu() {
	fmt.Print("1. Check Balance\n2. Deposit Money\n3. Withdraw Money\n4. Exit\n")
	fmt.Scan(&userSelectedOption)
}

func atmSystem(userInput int) {
	if !(validateUserInput(userInput)) {
		return
	}

	if userInput == 1 {
		showBalance()
		return
	}

	if userInput == 2 {
		depositMoney()
		return
	}

	if userInput == 3 {
		withdrawMoney()
		return
	}

	if userInput == 4 {
		fmt.Println("\nThank you for banking with us! Goodbye 👋")
		os.Exit(0)
	}
}

// FULL PROGRAM
func main() {
	for {
		showMenu()
		atmSystem(userSelectedOption)
	}
}
