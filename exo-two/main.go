// Exercise 2: Counting the Number of Characters

package main

import "fmt"

func main() {
	fmt.Print("What is the input string? ")
	var input string
	fmt.Scanln(&input)
	if len(input) == 0 {
		fmt.Print("error: the inputed string is empty!")
		return
	}
	fmt.Printf("%s has %d characters.", input, len(input))
}
