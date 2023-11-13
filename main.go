package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

const dailySalary = 500

type Day struct {
	IN      string
	OUT     string
	isRest  bool
	dayType string
}

func main() {
	days := make([]Day, 7)
	for i := range days {
		days[i] = Day{
			IN:      "0900",
			OUT:     "0900",
			isRest:  false,
			dayType: "Normal Day",
		}
	}

	for i := 5; i < 7; i++ {
		days[i].isRest = true
	}

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
			modifyConfiguration(days)
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

func calculateHours(inTime, outTime string) (int, int, int) {
	in, _ := time.Parse("1504", inTime)
	out, _ := time.Parse("1504", outTime)
	nightShiftStart := 22
	nightShiftEnd := 6

	overtimeStart := in.Add(9 * time.Hour)
	if overtimeStart.Before(in) {
		overtimeStart = overtimeStart.Add(24 * time.Hour)
	}

	overtimeHours := 0
	overtimeNightShiftHours := 0
	regularNightShiftHours := 0

	if out.Before(in) {
		out = out.Add(24 * time.Hour)
	}

	for t := in; t.Before(overtimeStart) && t.Before(out); t = t.Add(time.Hour) {
		if t.Hour() >= nightShiftStart || t.Hour() < nightShiftEnd {
			regularNightShiftHours++
		}
	}

	for t := overtimeStart; t.Before(out); t = t.Add(time.Hour) {
		if t.Hour() >= nightShiftStart || t.Hour() < nightShiftEnd {
			overtimeNightShiftHours++
		} else {
			overtimeHours++
		}
	}

	return overtimeHours, regularNightShiftHours, overtimeNightShiftHours
}

func isValidMilitaryTime(timeStr string) bool {
	militaryTimeRegex := regexp.MustCompile(`^([01]\d|2[0-3])([0-5]\d)$`)
	return militaryTimeRegex.MatchString(timeStr)
}

func isValidDayType(dayType string) bool {
	validDayTypes := map[string]bool{
		"Normal Day": true,
		"SNWH":       true,
		"RH":         true,
	}

	_, isValid := validDayTypes[dayType]
	return isValid
}

func modifyConfiguration(days []Day) {

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
