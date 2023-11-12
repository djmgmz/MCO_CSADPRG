package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("1) Generate Payroll")
		fmt.Println("2)	Modify Configuration")
		fmt.Println("3) Exit")

		var choice int
		fmt.Println("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			generatePayroll()
		case 2:
			modifyConfiguration()
		case 3:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid Choice. Please try again.")
		}
	}
}

func generatePayroll() {
	fmt.Println("Payroll Generated")
}

func modifyConfiguration() {
	fmt.Println("Modifying Configuration...")
}
