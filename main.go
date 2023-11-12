package main

import (
	"fmt"
	"os"
	"regexp"
)

const dailySalary = 500

type Day struct {
	IN      string
	OUT     string
	isRest  bool
	dayType string
}

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
			computePayroll()
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

func computePayroll() {

	fmt.Println("Payroll Generated")
}

func isValidMilitaryTime(timeStr string) bool {
	militaryTimeRegex := regexp.MustCompile(`^([01]\d|2[0-3])([0-5]\d)$`)
	return militaryTimeRegex.MatchString(timeStr)
}

func isValidDayType(dayType string) bool {
	validDayTypes := map[string]bool{
		"Normal Day":    true,
		"Rest Day":      true,
		"SNWH":          true,
		"SNWH-Rest Day": true,
		"RH":            true,
		"RH-Rest Day":   true,
	}

	_, isValid := validDayTypes[dayType]
	return isValid
}

func modifyConfiguration() {
	days := make([]Day, 7)

	for {
		for i := 1; i <= 7; i++ {
			fmt.Println(i, ")", i, "th Day")
		}

		var choice int
		fmt.Println("Enter your choice (0 to exit): ")
		fmt.Scan(&choice)

		if choice < 0 || choice > 7 {
			fmt.Println("Invalid choice. Please enter a number between 1 and 7, or 0 to exit.")
			continue
		}

		if choice == 0 {
			break
		}

		dayIndex := choice - 1

		selectedDay := &days[dayIndex]

		fmt.Printf("Selected %dth Day:\n", choice)
		fmt.Printf("IN: %s\nOUT: %s\nIsRest: %t\nDayType: %s\n", selectedDay.IN, selectedDay.OUT, selectedDay.isRest, selectedDay.dayType)

		for {
			fmt.Println("Select property to modify:")
			fmt.Println("1) IN")
			fmt.Println("2) OUT")
			fmt.Println("3) IsRest")
			fmt.Println("4) DayType")
			fmt.Println("0) Back to main menu")

			var subChoice int
			fmt.Scan(&subChoice)

			switch subChoice {
			case 0:
				break
			case 1:
				fmt.Printf("Current IN time: %s\n", selectedDay.IN)
				fmt.Println("Enter new IN time (in military time format HHmm):")
				var newIN string
				fmt.Scan(&newIN)
				if !isValidMilitaryTime(newIN) {
					fmt.Println("Invalid military time format. Please enter in HHmm format.")
					continue
				}
				selectedDay.IN = newIN
			case 2:
				fmt.Printf("Current OUT time: %s\n", selectedDay.OUT)
				fmt.Println("Enter new OUT time (in military time format HHmm):")
				var newOUT string
				fmt.Scan(&newOUT)
				if !isValidMilitaryTime(newOUT) {
					fmt.Println("Invalid military time format. Please enter in HHmm format.")
					continue
				}
				selectedDay.OUT = newOUT
			case 3:
				fmt.Printf("Current IsRest: %t\n", selectedDay.isRest)
				fmt.Println("Is it a rest day? (true/false):")
				var newIsRest string
				fmt.Scan(&newIsRest)
				if newIsRest == "true" {
					selectedDay.isRest = true
				} else if newIsRest == "false" {
					selectedDay.isRest = false
				} else {
					fmt.Println("Invalid input. Please enter true or false.")
					continue
				}
			case 4:
				fmt.Printf("Current DayType: %s\n", selectedDay.dayType)
				fmt.Println("Enter new day type:")
				var newDayType string
				fmt.Scan(&newDayType)
				if !isValidDayType(newDayType) {
					fmt.Println("Invalid day type. Please enter a valid day type.")
					continue
				}
				selectedDay.dayType = newDayType
			default:
				fmt.Println("Invalid choice. Please enter a number between 0 and 4.")
			}
		}
	}
	fmt.Println("Exiting Modify Configuration...")
}
