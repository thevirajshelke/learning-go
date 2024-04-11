package main

import "fmt"

// var helloworld string = "Hello World" // valid
var helloworld = "Hello World" // valid

// helloworld := "Hello World" // not valid

func main() {
	fmt.Println(helloworld)
	/*
		Declare a new variable & assing a value - initialization of a variable
		we can use ':=' to create a new variable & assign value to it
	*/
	// var card string = "Ace of Spades" // explicit type declaration
	card := "Ace of Spades" // implicit type assumption

	card = "Five of Diamonds" // re assigning a new value to the variable

	fmt.Println(card)
}
