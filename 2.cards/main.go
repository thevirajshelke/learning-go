package main

import "fmt"

func main() {
	/*
		Declare a new variable & assing a value - initialization of a variable
		we can use ':=' to create a new variable & assign value to it
	*/
	// var card string = "Ace of Spades" // explicit type declaration
	card := "Ace of Spades" // implicit type assumption

	card = "Five of Diamonds" // re assigning a new value to the variable

	fmt.Println(card)
}
