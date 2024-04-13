package main

import "fmt"

type deck []string

/*
	d is the receiver & print is a receiver function
	d is the actual reference of the varaible of type deck
	we generally use a convention in go to keep the receiver var of len 2-3 characters
	we can think of the var d as 'this' or 'self' from other programming languages but
	go is not a object oriented programming knowledge so 'this'. 'self' mean nothing in go and
	are not used anywhere in the language constructs.

	This is vaguely vaguely similar to the OO approach where a cards class would have had a 'cards'
	data member and print as the member function.
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
