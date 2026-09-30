package main

import "fmt"

func main() {
	//Income
	fmt.Printf("Hello, please enter your bi-weekly paycheck: ")
	fmt.Scanln(&biWeekly)

	//Expenses
	getExpensees()
	fmt.Printf(" \nYour monthly expenses are: $%d\nYour nmonthly income is: $%d\nLeftover funds: $%d\n", expenses, monthly, leftovers)
	anaylyzeExpenes()
	//Budget leftovers
	menu()

}
