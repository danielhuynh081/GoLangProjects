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

func anaylyzeExpenses() {
	//No more than 30%-35% of your income should go to rent
	//No more than 10%-15% of your income should go to utilities
	//no more than 10%-15% of your income should go to food
	//monthly icnoem should be ~50% for needs, ~30% wants ~20% investing and savings
	if rent > monthly*35/100 {
		fmt.Printf("\nYou are spending too much on rent! (~%d%% of your income)\nConsider finding a cheaper place to live.\n", rent*100/monthly)
	}
	if util > monthly*15/100 {
		fmt.Printf("\nYou are spending too much on utilities! Consider finding ways to reduce your utility bills.\n")
	}
	if food > monthly*15/100 {
		fmt.Printf("\nYou are spending too much on food! Consider finding ways to reduce your food expenses.\n")
	}
	if phone > monthly*2/100 {
		fmt.Printf("\nYou are spending too much on your phone! Consider finding a cheaper phone plan.\n")
	}
	if car > monthly*10/100 {
		fmt.Printf("\nYour car payments are too high, consider finding a cheaper car or refinancing your car loan.\n")
	}
}
