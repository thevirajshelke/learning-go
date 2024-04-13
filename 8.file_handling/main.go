package main

func main() {
	colorsslice := colors{"Red", "Blue", "Yellow"}
	colorsslice.print()
	colorsslice.saveToFile("colors")

	colorsslice2 := readColorsFromFile("colors")
	colorsslice2.print()
}
