package main

func main() {
	cards := newDeck()
	cards.print()

}

/*
This is a function declaration & definition . In go we have to explicitly define the return type
*/
func newCard() string {
	return "Five of Diamonds"
}
