// Exercise 1: Saying hello

package main

import "fmt"

func main() {
	fmt.Print("What is your name? ")
	var input string
	fmt.Scanln(&input)
	fmt.Printf("Hello, %s, nice to meet you!", input)
}
