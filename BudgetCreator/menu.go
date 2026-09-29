package main

import "fmt"

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
		frameowrk4()
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
