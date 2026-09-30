package main

import "fmt"

var biWeekly, rent, util, food, phone, gas, savings, investments, wants, expenses, monthly, leftovers, car int

func getExpensees() {
	monthly = biWeekly * 2

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

	fmt.Printf("How much of your paycheck goes to a car payment every month: ")
	fmt.Scanln(&car)

	//Count Expenses
	expenses += rent + util + food + phone + gas + car
	leftovers = monthly - expenses
}
