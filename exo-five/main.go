// Exercise 5: Simple Math

package main

import (
	"fmt"
	"os"
	"strconv"
)

func getNumber(name string) int64 {
	var raw string
	fmt.Printf("What is the %s number? ", name)
	fmt.Scanln(&raw)
	num, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		fmt.Printf("error with %s number: %s", name, err)
		os.Exit(1)
	}
	return num
}

func main() {
	operandOne := getNumber("first")
	operandTwo := getNumber("second")

	fmt.Printf("%d + %d = %d\n", operandOne, operandTwo, operandOne+operandTwo)
	fmt.Printf("%d - %d = %d\n", operandOne, operandTwo, operandOne-operandTwo)
	fmt.Printf("%d * %d = %d\n", operandOne, operandTwo, operandOne*operandTwo)
	fmt.Printf("%d / %d = %f\n", operandOne, operandTwo, float64(operandOne)/float64(operandTwo))
}
