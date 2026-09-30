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

// 30/30/30/10 Framework
func framework4() {
	necessities := leftovers * 20 / 100
	goals := leftovers * 70 / 100
	wants := leftovers * 10 / 100

	fmt.Printf("\nYour biweekly budget is as follows:\n Necessities/Emeregency: $%d\n Goals: $%d\n Wants: $%d\n",
		necessities/2, goals/2, wants/2)
	fmt.Printf("\nYour monthly budget is as follows:\n Necessities/Emergency: $%d\n Goals: $%d\n Wants: $%d\n",
		necessities, goals, wants)

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
	fmt.Printf("How much of your paycheck goes to a car payment every month: ")
	fmt.Scanln(&car)
	fmt.Printf("How much of your paycheck goes to gas every month: ")
	fmt.Scanln(&gas)
}

func anaylyzeExpenes() {
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
