// Exercise 8: Pizza Party

package main

import (
	"fmt"
	"math"
	"strconv"
)

const piecesPerPizza = 8

func main() {
	var rawPeople string
	fmt.Print("How many people? ")
	fmt.Scanln(&rawPeople)
	people, err := strconv.ParseInt(rawPeople, 10, 64)
	if err != nil {
		fmt.Printf("error with people amount: %s", err)
		return
	}

	var rawPizzas string
	fmt.Print("How many pizzas do you have? ")
	fmt.Scanln(&rawPizzas)
	pizzas, err := strconv.ParseInt(rawPizzas, 10, 64)
	if err != nil {
		fmt.Printf("error with pizzas amount: %s", err)
	}

	plural := ""
	if pizzas > 1 {
		plural = "s"
	}
	fmt.Printf("%d with %d pizza%s\n", people, pizzas, plural)

	pizzaSlices := pizzas * piecesPerPizza

	plural = ""
	piecePerPerson := int(math.Floor(float64(pizzaSlices) / float64(people)))
	if piecePerPerson > 1 {
		plural = "s"
	}

	fmt.Printf("Each person gets %d piece%s\n", piecePerPerson, plural)

	plural = ""
	leftovers := pizzaSlices % people
	if leftovers > 1 {
		plural = "s"
	}
	fmt.Printf("There are %d leftover piece%s\n", leftovers, plural)
}
