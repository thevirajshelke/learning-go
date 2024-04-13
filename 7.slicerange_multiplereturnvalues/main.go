package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println("Printing the new deck")
	cards.print()
	hand, remainingDeck := deal(cards, 5)
	fmt.Println("Printing the new hand")
	hand.print()
	fmt.Println("Printing the remaining deck")
	remainingDeck.print()
}

/*
This is a function declaration & definition . In go we have to explicitly define the return type
*/
func newCard() string {
	return "Five of Diamonds"
}
