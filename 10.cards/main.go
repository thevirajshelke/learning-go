package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println("Printing the new deck")
	cards.print()
	hand, remainingDeck := deal(cards, 5)
	hand.saveToFile("hand")
	remainingDeck.saveToFile("remainingDeck")
	fmt.Println("Printing the new hand from file")
	hand = readDeckFromFile("hand")
	hand.print()
	fmt.Println("Printing the remaining deck from file")
	remainingDeck = readDeckFromFile("remainingDeck")
	remainingDeck.print()
	fmt.Println("After shuffling!")
	cards.shuffle()
	cards.print()
}

/*
This is a function declaration & definition . In go we have to explicitly define the return type
*/
func newCard() string {
	return "Five of Diamonds"
}
