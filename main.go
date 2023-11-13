package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

const defaultSalary = 500

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
			computePayroll(days)
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

func computePayroll(days []Day) {
	weeklySalary := 0.0
	for i := 0; i < len(days); i++ {
		dailySalary := 0.0
		day := days[i]
		overtimeHours, regularNightShiftHours, overtimeNightShiftHours := calculateHours(day.IN, day.OUT)
		fmt.Printf("Day %d:\n", i+1)

		if day.IN == day.OUT && day.isRest == false {
			fmt.Println("ABSENT")
		} else if day.IN == day.OUT && day.isRest == true {
			fmt.Println("REST DAY")
		} else {
			fmt.Printf(" IN: %s\n OUT: %s\n DayType: %s\n IsRest: %t\n", day.IN, day.OUT, day.dayType, day.isRest)
			if day.dayType == "Normal Day" && day.isRest == false {
				fmt.Printf("Base Salary: %.2f\n", float64(defaultSalary))
				dailySalary += defaultSalary

				OT := float64(overtimeHours) * 500 / 8 * 1.25
				NS := float64(regularNightShiftHours) * 500 / 8 * 1.10
				OTNS := float64(overtimeNightShiftHours) * 500 / 8 * 1.375

				if overtimeHours > 0 {
					fmt.Printf("Hours OT x OT Hourly Rate\n")
					fmt.Printf("= %d x 500 / 8 * 1.25\n", overtimeHours)
					fmt.Printf("= %.2f\n", OT)
				}

				if regularNightShiftHours > 0 {
					fmt.Printf("Hours on NS x Hourly Rate x NSD\n")
					fmt.Printf("= %d x 500 / 8 * 1.10\n", regularNightShiftHours)
					fmt.Printf("= %.2f\n", NS)
				}

				if overtimeNightShiftHours > 0 {
					fmt.Printf("= %d x 500 / 8 * 1.375\n", overtimeNightShiftHours)
					fmt.Printf("= %.2f\n", OTNS)
				}
				dailySalary += OT + NS + OTNS
			} else if day.dayType == "Normal Day" && day.isRest == true {
				fmt.Printf("Base Salary: %.2f\n", float64(defaultSalary*1.3))
				dailySalary += defaultSalary * 1.3

				OT := float64(overtimeHours) * 500 / 8 * 1.69
				NS := float64(regularNightShiftHours) * 500 / 8 * 1.10
				OTNS := float64(overtimeNightShiftHours) * 500 / 8 * 1.859

				if overtimeHours > 0 {
					fmt.Printf("Hours OT x OT Hourly Rate\n")
					fmt.Printf("= %d x 500 / 8 * 1.69\n", overtimeHours)
					fmt.Printf("= %.2f\n", OT)
				}

				if regularNightShiftHours > 0 {
					fmt.Printf("Hours on NS x Hourly Rate x NSD\n")
					fmt.Printf("= %d x 500 / 8 * 1.10\n", regularNightShiftHours)
					fmt.Printf("= %.2f\n", NS)
				}

				if overtimeNightShiftHours > 0 {
					fmt.Printf("= %d x 500 / 8 * 1.859\n", overtimeNightShiftHours)
					fmt.Printf("= %.2f\n", OTNS)
				}
				dailySalary += OT + NS + OTNS
			} else if day.dayType == "SNWH" && day.isRest == false {
				fmt.Printf("Base Salary: %.2f\n", float64(defaultSalary*1.3))
				dailySalary += defaultSalary * 1.3

				OT := float64(overtimeHours) * 500 / 8 * 1.69
				NS := float64(regularNightShiftHours) * 500 / 8 * 1.10
				OTNS := float64(overtimeNightShiftHours) * 500 / 8 * 1.859
				if overtimeHours > 0 {
					fmt.Printf("Hours OT x OT Hourly Rate\n")
					fmt.Printf("= %d x 500 / 8 * 1.69\n", overtimeHours)
					fmt.Printf("= %.2f\n", OT)
				}

				if regularNightShiftHours > 0 {
					fmt.Printf("Hours on NS x Hourly Rate x NSD\n")
					fmt.Printf("= %d x 500 / 8 * 1.10\n", regularNightShiftHours)
					fmt.Printf("= %.2f\n", NS)
				}

				if overtimeNightShiftHours > 0 {
					fmt.Printf("= %d x 500 / 8 * 1.859\n", overtimeNightShiftHours)
					fmt.Printf("= %.2f\n", OTNS)
				}
				dailySalary += OT + NS + OTNS
			} else if day.dayType == "SNWH" && day.isRest == true {
				fmt.Printf("Base Salary: %.2f\n", float64(defaultSalary*1.5))
				dailySalary += defaultSalary * 1.5

				OT := float64(overtimeHours) * 500 / 8 * 1.95
				NS := float64(regularNightShiftHours) * 500 / 8 * 1.10
				OTNS := float64(overtimeNightShiftHours) * 500 / 8 * 2.145
				if overtimeHours > 0 {
					fmt.Printf("Hours OT x OT Hourly Rate\n")
					fmt.Printf("= %d x 500 / 8 * 1.95\n", overtimeHours)
					fmt.Printf("= %.2f\n", OT)
				}

				if regularNightShiftHours > 0 {
					fmt.Printf("Hours on NS x Hourly Rate x NSD\n")
					fmt.Printf("= %d x 500 / 8 * 1.10\n", regularNightShiftHours)
					fmt.Printf("= %.2f\n", NS)
				}

				if overtimeNightShiftHours > 0 {
					fmt.Printf("= %d x 500 / 8 * 2.145\n", overtimeNightShiftHours)
					fmt.Printf("= %.2f\n", OTNS)
				}
				dailySalary += OT + NS + OTNS
			} else if day.dayType == "RH" && day.isRest == false {
				fmt.Printf("Base Salary: %.2f\n", float64(defaultSalary*2))
				dailySalary += defaultSalary * 2

				OT := float64(overtimeHours) * 500 / 8 * 2.6
				NS := float64(regularNightShiftHours) * 500 / 8 * 1.10
				OTNS := float64(overtimeNightShiftHours) * 500 / 8 * 2.86
				if overtimeHours > 0 {
					fmt.Printf("Hours OT x OT Hourly Rate\n")
					fmt.Printf("= %d x 500 / 8 * 2.6\n", overtimeHours)
					fmt.Printf("= %.2f\n", OT)
				}

				if regularNightShiftHours > 0 {
					fmt.Printf("Hours on NS x Hourly Rate x NSD\n")
					fmt.Printf("= %d x 500 / 8 * 1.10\n", regularNightShiftHours)
					fmt.Printf("= %.2f\n", NS)
				}

				if overtimeNightShiftHours > 0 {
					fmt.Printf("= %d x 500 / 8 * 2.86\n", overtimeNightShiftHours)
					fmt.Printf("= %.2f\n", OTNS)
				}
				dailySalary += OT + NS + OTNS
			} else if day.dayType == "RH" && day.isRest == true {
				fmt.Printf("Base Salary: %.2f\n", float64(defaultSalary*2.6))
				dailySalary += defaultSalary * 2.6

				OT := float64(overtimeHours) * 500 / 8 * 3.38
				NS := float64(regularNightShiftHours) * 500 / 8 * 1.10
				OTNS := float64(overtimeNightShiftHours) * 500 / 8 * 3.718
				if overtimeHours > 0 {
					fmt.Printf("Hours OT x OT Hourly Rate\n")
					fmt.Printf("= %d x 500 / 8 * 3.38\n", overtimeHours)
					fmt.Printf("= %.2f\n", OT)
				}

				if regularNightShiftHours > 0 {
					fmt.Printf("Hours on NS x Hourly Rate x NSD\n")
					fmt.Printf("= %d x 500 / 8 * 1.10\n", regularNightShiftHours)
					fmt.Printf("= %.2f\n", NS)
				}

				if overtimeNightShiftHours > 0 {
					fmt.Printf("= %d x 500 / 8 * 3.718\n", overtimeNightShiftHours)
					fmt.Printf("= %.2f\n", OTNS)
				}
				dailySalary += OT + NS + OTNS
			}
			fmt.Printf("Daily Salary = %.2f", dailySalary)
			fmt.Println("")
		}

		fmt.Println("")
		weeklySalary += dailySalary
	}
	fmt.Printf("Weekly Salary: %.2f\n", weeklySalary)
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

	outOfMenu:
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
				break outOfMenu
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
