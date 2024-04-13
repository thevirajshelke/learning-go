package main

func main() {

	/*
		cards is of type 'deck' which is declared in deck.go
		'deck' is '[]string'. symbols in same package need not be exported and are available to access
	*/
	cards := deck{"Ace of Diamonds", getCard()}
	/*
		'print' is a receiver function avaiable on type 'deck'
		since cards is of type 'deck' it gets access to 'print' method
	*/
	cards.print()

}

func getCard() string {
	return "Ace of Hearts"
}
