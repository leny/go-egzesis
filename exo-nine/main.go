// Exercise 9: Paint Calculator

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func getNumber(name string) int64 {
	var raw string
	fmt.Printf("What is the %s of the room, in feet? ", name)
	fmt.Scanln(&raw)
	num, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		fmt.Printf("error with %s number: %s", name, err)
		os.Exit(1)
	}
	return num
}

func getDimension() (int64, int64) {
	length := getNumber("length")
	width := getNumber("width")
	return length, width
}

const gallonPerSquareFeet = 350.0

func main() {
	length, width := getDimension()
	area := length * width
	gallons := int(math.Ceil(float64(area) / gallonPerSquareFeet))
	fmt.Printf("You will need to purchase %d gallons of paint to cover %d square feet.\n", gallons, area)
}
