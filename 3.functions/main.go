package main

import "fmt"

func main() {
	card := newCard() // go will infer the datatype of 'card' by looking at return type of 'newCard()'
	fmt.Println(card)
}

/*
This is a function declarationn & definition . In go we have to explicitly define the return type
*/
func newCard() string {
	return "Five of Diamonds"
}
