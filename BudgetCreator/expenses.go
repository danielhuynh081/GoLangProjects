package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Budget struct {
	biWeekly    int
	rent        int
	util        int
	food        int
	phone       int
	gas         int
	car         int
	savings     int
	investments int
	wants       int
	expenses    int
	monthly     int
	leftovers   int
}

var b Budget

func readInt(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}
		input = strings.TrimSpace(input)
		value, err := strconv.Atoi(input)
		if err != nil || value < 0 {
			fmt.Println("Invalid input. Please enter a non-negative integer.")
			continue
		}
		return value
	}
}

func getExpensees() {
	b.monthly = b.biWeekly * 2

	b.rent = readInt("How much of your paycheck goes to rent every month: ")
	b.util = readInt("How much of your paycheck goes to utilities every month: ")
	b.food = readInt("How much of your paycheck goes to food every month: ")
	b.phone = readInt("How much of your paycheck goes to phone every month: ")
	b.gas = readInt("How much of your paycheck goes to gas every month: ")
	b.car = readInt("How much of your paycheck goes to a car payment every month: ")

	b.expenses = b.rent + b.util + b.food + b.phone + b.gas + b.car
	b.leftovers = b.monthly - b.expenses
}

func anaylyzeExpenses() {
	if b.monthly == 0 {
		fmt.Println("Monthly income is zero. Cannot perform analysis.")
		return
	}

	fmt.Println("\n--- Expense Analysis ---")

	printAnalysis("Rent", b.rent, b.monthly, 35)
	printAnalysis("Utilities", b.util, b.monthly, 15)
	printAnalysis("Food", b.food, b.monthly, 15)
	printAnalysis("Phone", b.phone, b.monthly, 2)
	printAnalysis("Car", b.car, b.monthly, 10)

	expenseRatio := (b.expenses * 100) / b.monthly
	fmt.Printf("\nTotal Expense Ratio: %d%% of monthly income.\n", expenseRatio)

	if b.leftovers < 0 {
		fmt.Printf("Warning: You are in a deficit of $%d per month.\n", -b.leftovers)
	} else {
		fmt.Printf("Remaining funds for savings, investments, and leisure: $%d\n", b.leftovers)
	}
}

func printAnalysis(category string, amount, total, threshold int) {
	percentage := (amount * 100) / total
	fmt.Printf("%s: $%d (%d%% of income)\n", category, amount, percentage)
	if percentage > threshold {
		fmt.Printf("  - Recommendation: %s exceeds the recommended %d%% threshold. Consider reducing this expense.\n", category, threshold)
	}
}
