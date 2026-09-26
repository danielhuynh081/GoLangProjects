package main

import "fmt"

func framework1() {
	//50/30/20 framework
	savings = leftovers * 20 / 100
	investments = leftovers * 30 / 100
	wants = leftovers * 50 / 100

	fmt.Printf("\nYour monthly budget is as follows:\n Savings: $%d\n Investments: $%d\n Wants: $%d\n", savings, investments, wants)
	fmt.Printf("\nYour biweekly budget is as follows:\n Savings: $%d\n Investments: $%d\n Wants: $%d\n", savings/2, investments/2, wants/2)
}

// Zero-Based Budgeting
func framework2() {
	// Calculate remainder after expenses

	if leftovers < 0 {
		fmt.Printf("\nYou went over your monthly income by $%d!\n", -leftovers)
		return
	}

	wants = leftovers * 20 / 100
	leftovers = leftovers - wants

	fmt.Printf("\nYour biweekly budget is as follows:\n Rent: $%d\n Utilities: $%d\n Food: $%d\n Phone: $%d\n Gas: $%d\n Investments/Savings: $%d\n Wants: $%d\n",
		rent/2, util/2, food/2, phone/2, gas/2, leftovers/2, wants/2)
	fmt.Printf("\nYour monthly expenses is as follows:\n Rent: $%d\n Utilities: $%d\n Food: $%d\n Phone: $%d\n Gas: $%d\n Investments/Savings: $%d\n Wants: $%d\n",
		rent, util, food, phone, gas, leftovers, wants)

}

// Pay-Yourself-First
func framework3() {
	if leftovers < 0 {
		fmt.Printf("\nYou went over your monthly income by $%d!\n", -leftovers)
		return
	}

	wants = leftovers * 10 / 100
	leftovers = leftovers - wants
	debt := leftovers * 80 / 100
	savings := leftovers * 10 / 100

	fmt.Printf("\nYour biweekly budget is as follows:\n Rent: $%d\n Utilities: $%d\n Food: $%d\n Phone: $%d\n Gas: $%d\n Debt: $%d\n Investments/Savings: $%d\n Wants: $%d\n",
		rent/2, util/2, food/2, phone/2, gas/2, debt/2, savings/2, wants/2)
	fmt.Printf("\nYour monthly budget is as follows:\n Rent: $%d\n Utilities: $%d\n Food: $%d\n Phone: $%d\n Gas: $%d\n Debt: $%d\n Investments/Savings: $%d\n Wants: $%d\n",
		rent, util, food, phone, gas, debt, savings, wants)
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
