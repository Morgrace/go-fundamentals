package main

import (
	"fmt"
	"time"
)

func addExpenseToCategory(expense map[string]int, key string) {
	fmt.Print("Enter amount($):")
	fmt.Scan(&expenseAmount)
	if !(expenseAmount > 0) {
		fmt.Print("\n🛑 Expense amount must be greater than zero (0)\n")
		return
	}
	if !(totalBalance > expenseAmount) {
		fmt.Println("\n🛑 Expense is greater than available balance: Please consider purchases something less expensive")
		return
	}
	expenseCategory = key
	totalBalance -= expenseAmount
	expense[key] += expenseAmount
	fmt.Printf("\n🥳 Successfully added %v to your expenses\n\n - - -\n\nYou have 💰 $%v left in your wallet\n\n", key, totalBalance)
	time.Sleep(time.Second * 3)
}

func addExpense() {
	fmt.Println("\n<<Select Category>>\n1. Food\n2. Transport\n3. Entertainment\n4. Utilities\n5. Others")
	fmt.Scan(&userSelection)

	switch userSelection {
	case 1:
		addExpenseToCategory(expenses, "food")
	case 2:
		addExpenseToCategory(expenses, "transport")
	case 3:
		addExpenseToCategory(expenses, "entertainment")
	case 4:
		addExpenseToCategory(expenses, "utilities")
	case 5:
		addExpenseToCategory(expenses, "others")
	default:
		fmt.Println("Invalid category. Please try again")
	}
}
