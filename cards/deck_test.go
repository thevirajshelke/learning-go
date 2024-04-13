package main

import (
	"os"
	"testing"
)

// use 'go mod init cards' for running tests

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("The new deck length should have been 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("The new deck should have started with 'Ace of Spades' but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("The new deck should have ended with 'Four of Clubs' but got %v", d[0])
	}
}

func TestDeckFileSavingAndReading(t *testing.T) {
	os.Remove("_testingdeck")
	d := newDeck()
	d.saveToFile("_testingdeck")

	loadedDeck := readDeckFromFile("_testingdeck")

	if len(loadedDeck) != 16 {
		t.Errorf("The new deck length should have been 16 but got %v", len(d))
	}

}
