package main

import "fmt"

func main() {
	fmt.Println("Welcome to BudgetMe!")
	
	b.biWeekly = readInt("Please enter your bi-weekly paycheck: ")
	
	getExpenses()
	
	fmt.Printf("\nYour monthly expenses are: $%d\n", b.expenses)
	fmt.Printf("Your monthly income is: $%d\n", b.monthly)
	fmt.Printf("Leftover funds: $%d\n", b.leftovers)
	
	analyzeExpenses()

	menu()
}
