// Exercise 3: Printing Quotes

package main

import "fmt"

func main() {
	var quote string
	fmt.Print("What is the quote? ")
	fmt.Scanln(&quote)
	if len(quote) == 0 {
		fmt.Print("error: the quote cannot be empty!")
		return
	}

	var author string
	fmt.Print("Who said it? ")
	fmt.Scanln(&author)
	if len(author) == 0 {
		author = "Unknown"
	}

	fmt.Printf("%s says, \"%s\".", author, quote)
}
