// Exercise 6: Retirement Calculator

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var rawAge string
	fmt.Print("What is your current age? ")
	fmt.Scanln(&rawAge)
	age, err := strconv.ParseInt(rawAge, 10, 16)
	if err != nil {
		fmt.Printf("error with age: %s\n", err)
	}

	var rawPension string
	fmt.Print("At what age would you like to retire? ")
	fmt.Scanln(&rawPension)
	pension, err := strconv.ParseInt(rawPension, 10, 16)
	if err != nil {
		fmt.Printf("error with retire age: %s\n", err)
	}

	currentYear := time.Now().Year()
	yearsToWork := pension - age
	retireYear := currentYear + int(yearsToWork)

	if yearsToWork <= 0 {
		fmt.Print("You already can retire!\n")
		return
	}

	fmt.Printf("You have %d years left until you can retire.\nIt's %d, so you can retire in %d\n", yearsToWork, currentYear, retireYear)
}
