package main

import (
	"fmt"
	"os"
	"strings"
)

type colors []string

func (c colors) print() {
	for i, color := range c {
		fmt.Println(i, color)
	}
}

func (c colors) toString() string {
	// here we are converting colors of type []string into string
	return strings.Join([]string(c), ",")
}

func (c colors) saveToFile(filename string) error {
	// here we are converting string to []byte which is expected by write file function
	// 0666 is r/w for all
	return os.WriteFile(filename, []byte(c.toString()), 0666)
}

func readColorsFromFile(filename string) colors {
	bs, err := os.ReadFile(filename) // returns data & error - multiple return values
	// common error handling practice in go
	if err != nil {
		// option 1 - can we return a default colors slice with some data? or blank slice?
		// option 2 - can we fail the program by logging the error
		fmt.Println("Error:", err)
		os.Exit(1) // 1 indicates unsuccessful & 0 indicates success
	}
	// here we are conveting []byte to string & splitting with separator
	cs := strings.Split(string(bs), ",")
	// converting []string to colors so that we can access the receiver functions defined on type color
	return colors(cs)
}
