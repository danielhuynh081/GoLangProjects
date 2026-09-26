package main

import "fmt"

var biWeekly, rent, util, food, phone, gas, savings, investments, wants, expenses, monthly, leftovers int

func getExpensees() {
	monthly := biWeekly * 2

	fmt.Printf("How much of your paycheck goes to rent every month: ")
	fmt.Scanln(&rent)
	expenses += rent

	fmt.Printf("How much of your paycheck goes to utilities every month: ")
	fmt.Scanln(&util)
	expenses += util

	fmt.Printf("How much of your paycheck goes to food every month: ")
	fmt.Scanln(&food)
	expenses += food

	fmt.Printf("How much of your paycheck goes to phone every month: ")
	fmt.Scanln(&phone)
	expenses += phone

	fmt.Printf("How much of your paycheck goes to gas every month: ")
	fmt.Scanln(&gas)
	expenses += gas

	leftovers = monthly - expenses
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
		framework2()
	case 3:
		fmt.Println("You chose Pay-Yourself-First.")
		framework3()
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
