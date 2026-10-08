package main

import "fmt"

func framework1() {
	//50/30/20 framework
	b.savings = b.leftovers * 20 / 100
	b.investments = b.leftovers * 30 / 100
	b.wants = b.leftovers * 50 / 100

	fmt.Printf("\nYour monthly budget is as follows:\n Savings: $%d\n Investments: $%d\n Wants: $%d\n", b.savings, b.investments, b.wants)
	fmt.Printf("\nYour biweekly budget is as follows:\n Savings: $%d\n Investments: $%d\n Wants: $%d\n", b.savings/2, b.investments/2, b.wants/2)
}

// Zero-Based Budgeting
func framework2() {
	// Calculate remainder after expenses

	if b.leftovers < 0 {
		fmt.Printf("\nYou went over your monthly income by $%d!\n", -b.leftovers)
		return
	}

	b.wants = b.leftovers * 20 / 100
	remaining := b.leftovers - b.wants

	fmt.Printf("\nYour biweekly budget is as follows:\n Rent: $%d\n Utilities: $%d\n Food: $%d\n Phone: $%d\n Gas: $%d\n Investments/Savings: $%d\n Wants: $%d\n",
		b.rent/2, b.util/2, b.food/2, b.phone/2, b.gas/2, remaining/2, b.wants/2)
	fmt.Printf("\nYour monthly expenses is as follows:\n Rent: $%d\n Utilities: $%d\n Food: $%d\n Phone: $%d\n Gas: $%d\n Investments/Savings: $%d\n Wants: $%d\n",
		b.rent, b.util, b.food, b.phone, b.gas, remaining, b.wants)

}

// Pay-Yourself-First
func framework3() {
	if b.leftovers < 0 {
		fmt.Printf("\nYou went over your monthly income by $%d!\n", -b.leftovers)
		return
	}

	b.wants = b.leftovers * 10 / 100
	remaining := b.leftovers - b.wants
	debt := remaining * 80 / 100
	savings := remaining * 10 / 100

	fmt.Printf("\nYour biweekly budget is as follows:\n Rent: $%d\n Utilities: $%d\n Food: $%d\n Phone: $%d\n Gas: $%d\n Debt: $%d\n Investments/Savings: $%d\n Wants: $%d\n",
		b.rent/2, b.util/2, b.food/2, b.phone/2, b.gas/2, debt/2, savings/2, b.wants/2)
	fmt.Printf("\nYour monthly budget is as follows:\n Rent: $%d\n Utilities: $%d\n Food: $%d\n Phone: $%d\n Gas: $%d\n Debt: $%d\n Investments/Savings: $%d\n Wants: $%d\n",
		b.rent, b.util, b.food, b.phone, b.gas, debt, savings, b.wants)
}

// 30/30/30/10 Framework
func framework4() {
	necessities := b.leftovers * 20 / 100
	goals := b.leftovers * 70 / 100
	wants := b.leftovers * 10 / 100

	fmt.Printf("\nYour biweekly budget is as follows:\n Necessities/Emeregency: $%d\n Goals: $%d\n Wants: $%d\n",
		necessities/2, goals/2, wants/2)
	fmt.Printf("\nYour monthly budget is as follows:\n Necessities/Emergency: $%d\n Goals: $%d\n Wants: $%d\n",
		necessities, goals, wants)

}

func customizeBudget() {
	//Info gathering
	b.rent = readInt("How much of your paycheck goes to rent every month: ")
	b.util = readInt("How much of your paycheck goes to utilities every month: ")
	b.food = readInt("How much of your paycheck goes to food every month: ")
	b.phone = readInt("How much of your paycheck goes to phone every month: ")
	b.car = readInt("How much of your paycheck goes to a car payment every month: ")
	b.gas = readInt("How much of your paycheck goes to gas every month: ")

	b.expenses = b.rent + b.util + b.food + b.phone + b.gas + b.car
	b.leftovers = b.monthly - b.expenses
	
	fmt.Printf("\nNew total monthly expenses: $%d\n", b.expenses)
	fmt.Printf("New leftover funds: $%d\n", b.leftovers)
	analyzeExpenses()
}
