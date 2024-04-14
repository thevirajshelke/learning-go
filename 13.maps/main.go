package main

import "fmt"

func main() {
	/*
		Maps in go are smililar to hash in ruby, maps in javascript, dict in python
		The main difference between structs and maps in go is that maps need to define the key value type
		and actual value type. Basically all the keys in the maps must be same type and all values of same type
	*/
	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
	}

	var colors2 map[string]string
	colors3 := make(map[string]string)

	fmt.Println(colors)
	fmt.Println(colors2)
	fmt.Println(colors3)

	printMap(colors)

	colors3["black"] = "#000000"

	fmt.Println(colors3)
	fmt.Println(colors3["black"])

	delete(colors, "red")
	fmt.Println(colors["red"])
	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
