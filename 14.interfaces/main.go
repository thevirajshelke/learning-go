package main

import "fmt"

type bot interface {
	getGreeting() string
}

/*
	Since englishBot and marathiBot have a receiver function 'getGreeting' which returns a string
	they automatically also become of type 'bot'
*/
type englishBot struct{}
type marathiBot struct{}

func main() {
	eb := englishBot{}
	mb := marathiBot{}

	// eb & mb are of type bot as well! hence we can pass them in this function without error
	printGreeting(eb)
	printGreeting(mb)
}

// Following are identical functions!
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(mb marathiBot) {
// 	fmt.Println(mb.getGreeting())
// }

// Here b can take marathiBot as well as englishBot
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// very custom logic to generate english greeting
	return "Hi there"
}

// If we don't want to use the received value we can just ignore that variable and write the type as follows
func (marathiBot) getGreeting() string {
	// very custom logic to generate english greeting
	return "Namaskar"
}
