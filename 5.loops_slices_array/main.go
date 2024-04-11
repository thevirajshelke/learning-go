package main

import "fmt"

/*
In go we have two things,
1. array - fixed length, same type of values
2. slice - dynamic length (can grow and shrink), same type of values
*/

func main() {
	// Here we are declaring a slice of type strings and assigning 2 values to it
	cards := []string{"Six of Diamonds", getCard()}
	// Here we are adding a new value using append function.
	// Note that append actually creates a new slice with the value passed to it and returns that new slice
	// The old slice is not modified
	cards = append(cards, "Eight of Hearts")

	// Slices can be directly printed!
	fmt.Println(cards)

	/*
		'range' helps to iterate over the slice and in for loops we use two variables to catch the
		index and the actual value. Note ':=' is used on purpose because in go for loops the old varaibles
		are discarded hence we declare and initialize new vars using 'for i, cards := range cards'
	*/
	for i, card := range cards {
		// if we declare 'i' we need to use it or it's an error
		fmt.Println(i, card)
	}
}

func getCard() string {
	return "Five of Spades"
}
