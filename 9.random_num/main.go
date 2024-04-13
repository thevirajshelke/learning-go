package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	num := r.Intn(10)
	fmt.Println("Random number is", num)
	num2 := rand.Intn(10)
	fmt.Println("Random number is with default seeding", num2)

}
