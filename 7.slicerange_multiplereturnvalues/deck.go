package main

import "fmt"

type deck []string

/*
	This function returns multiple values
*/
func deal(d deck, handSize int) (deck, deck) {
	// The following syntax generates a new slice where the original slice is not modified
	/*
		slice[startingIndexIncluding:upToNotIncluding]
	*/
	return d[:handSize], d[handSize:]
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// If we don't want to use a variable put '_' this ignores the variable
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
