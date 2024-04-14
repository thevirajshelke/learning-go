package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

/*
	Note - Go is pass by value.
	The concept of slice and array is same as that of C & C++ i.e. the actual slice/array varaible name represents the pointer to the actual value.
	So when we add a receiver function to the slice or array and pass the slice or array any modification to this slice array will result into
	updation of the original varaible. Here go passes the slice or array pointer as a value & go being a pass by value remains true
	but this new value is just another pointer to same array or slice :)

	Some more types apart from slice & array that are references to the actual value,
	Following are reference types (No need of pointers in receiver functions)
	- slice/array
	- maps
	- channels
	- pointers
	- functions
	Following are value types (use pointers to update these in receiver function)
	- int
	- float
	- string
	- bool
	- structs

	When we create a slice, Go will automatically create which two data structures?
	An array and a structure that records the length of the slice, the capacity of the slice, and a reference to the underlying array

	When the slice is created, we get both an array and a data structure describing that array. The data structure describing the array *is* copied before being passed off to the function, but the underlying array is not.
*/
func main() {
	viraj := person{"Viraj", "Shelke"}
	viraj.print()
	viraj.updateFirstNameByValue("VIRAJJ")
	viraj.print()
	// & operator to get the memory address of the variable
	(&viraj).updateFirstNameByRef("VIRAJJ")
	viraj.print()
	// Go lets us do the following because it internally understands and manages the conversion that the receiver function is expecting
	viraj.updateFirstNameByRef("VIRAJJ2")
	viraj.print()
}

/*
	This is call by reference i.e. we are passing the address of the original structure
*/
// Note - A * with a type indicates a type of pointer and * with a variable is actually a dereferencing operator
func (ptr *person) updateFirstNameByRef(newFirstName string) {
	// * operator will dereference the pointer i.e. access the memory address
	(*ptr).firstName = newFirstName
}

// The following method is ineffective
// GO IS A PASS BY VALUE LANGUAGE!
// So the following is ineffective
func (p person) updateFirstNameByValue(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
