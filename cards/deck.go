package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) shuffle() {
	// seeding
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newposition := r.Intn(len(d) - 1)

		// the rhs is completely evaluated first and then assignment happens
		d[i], d[newposition] = d[newposition], d[i]
	}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func readDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	ds := strings.Split(string(bs), ",")
	return deck(ds)
}
