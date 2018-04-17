// Exercise 7: Area of a Rectangular Room

package main

import (
	"fmt"
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

const feetToMeters = 0.09290304

func main() {
	length, width := getDimension()
	fmt.Printf("You entered dimensions of %d feet by %d feet.\n", length, width)
	feetArea := length * width
	metersArea := float64(feetArea) * feetToMeters
	fmt.Printf("The area is\n%d square feet\n%f square meters", feetArea, metersArea)
}
