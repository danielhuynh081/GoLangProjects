package main

import "fmt"

var biWeekly, rent, util, food, phone, gas, savings, investments, wants, expenses, monthly, leftovers int

func getExpensees() {
	monthly = biWeekly * 2

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
