package main

import "fmt"

var biWeekly, rent, util, food, phone, gas, savings, investments, wants, buffer int

func main() {
	fmt.Printf("Hello, please enter your bi-weekly paycheck: ")
	fmt.Scanln(&biWeekly)

	//Income
	var monthly = biWeekly * 2
	fmt.Printf(" Your bi-weekly paycheck is: $%d\n monthly income: $%d\n", biWeekly, monthly)

	//Menu Option
	menu()

}

func menu() {
	var choice int

	fmt.Print(`
(1) 50/30/20 Framework (Currently comfortable, just need to save and invest more.)
(2) Zero-Based Budgeting (Currently struggling to make ends meet, need to budget every dollar.)
(3) Pay-Yourself-First (Currently in debt, need to focus on paying off debt first.)
(4) 30/30/30/10 Framework (Currently comfortable, but want to focus on saving and investing more and less wants.)
(5) Customize Budget
(6) Exit

Enter your choice: `)

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("You chose the 50/30/20 framework.")
		framework1()
	case 2:
		fmt.Println("You chose Zero-Based Budgeting.")
	case 3:
		fmt.Println("You chose Pay-Yourself-First.")
	case 4:
		fmt.Println("You chose the 30/30/30/10 Framework.")
	case 5:
		fmt.Println("You chose to customize your budget.")
		customizeBudget()
	case 6:
		fmt.Println("Goodbye!")
		return
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}

func framework1() {
	//50/30/20 framework
	savings = biWeekly * 2 * 20 / 100
	investments = biWeekly * 2 * 30 / 100
	wants = biWeekly * 2 * 50 / 100

	fmt.Printf("\nYour monthly budget is as follows:\n Savings: $%d\n Investments: $%d\n Wants: $%d\n", savings, investments, wants)
	fmt.Printf("\nYour biweekly budget is as follows:\n Savings: $%d\n Investments: $%d\n Wants: $%d\n", savings/2, investments/2, wants/2)
}

func customizeBudget() {
	//Info gathering

	fmt.Printf("How much of your paycheck goes to rent every month: ")
	fmt.Scanln(&rent)
	fmt.Printf("How much of your paycheck goes to utilities every month: ")
	fmt.Scanln(&util)
	fmt.Printf("How much of your paycheck goes to food every month: ")
	fmt.Scanln(&food)
	fmt.Printf("How much of your paycheck goes to phone every month: ")
	fmt.Scanln(&phone)
	fmt.Printf("How much of your paycheck goes to gas every month: ")
	fmt.Scanln(&gas)

}
