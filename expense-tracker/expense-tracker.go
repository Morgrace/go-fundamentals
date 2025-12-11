package main

import (
	"fmt"
	"os"
	"time"
)

var totalBalance = 1000

var expenses = make(map[string]int)

// User Input
var userSelection int
var expenseAmount int
var expenseCategory string

func showExpenses() {
	fmt.Printf("Your expenses are:\n🥫 Food: %v\n\n✈️ Transport: %v\n\n🥳 Entertainment: %v\n\n🚰 Utilities %v\n\n✳️ Others: %v\n\n", expenses["food"], expenses["transport"], expenses["entertainment"], expenses["utilities"], expenses["others"])
	time.Sleep(time.Second * 2)
}

func showTotalExpenses() {
	sum := 0
	for _, value := range expenses {
		sum += value
	}
	fmt.Printf("\n💰 Total expenses: $%v\n\n", sum)
	time.Sleep(time.Second * 2)
}
func deleteLastExpense() {
	if !(expenseAmount > 0) && len(expenseCategory) == 0 {
		fmt.Printf("\n🛑 You have not made any operation in this session! Add expenses to begin!\n")
		return
	}
	// Undo last operation
	expenses[expenseCategory] = expenses[expenseCategory] - expenseAmount
	totalBalance += expenseAmount
	// Reset variables
	expenseCategory = ""
	expenseAmount = 0

}

func validateMenuInput() {
	fmt.Scan(&userSelection)
	switch userSelection {
	case 1:
		addExpense()
	case 2:
		showExpenses()
	case 3:
		showTotalExpenses()
	case 4:
		fmt.Printf("\n💰 Balance: $%v\n\n", totalBalance)
		time.Sleep(time.Second * 2)
	case 5:
		deleteLastExpense()
		time.Sleep(time.Second * 2)

	case 6:
		os.Exit(0)
	default:
		fmt.Println("\nInvalid option. Please try again")
	}
}

func showMenu() {
	fmt.Println("\n<<<<<<Main Menu>>>>>\n1. Add Expense\n2. View Expenses by Category\n3. View Total Expenses\n4. View Remaining Balance\n5. Delete Last Expense\n6. Exit")
}
func main() {
	for {
		showMenu()
		validateMenuInput()
	}
}
