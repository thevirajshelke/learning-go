package main

import "fmt"

type optionalInfo struct {
	nickname string
}

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	optionalInfo
}

func main() {
	// first value is assigned to first field - not recommended because this might cause errors if the struct is modified later
	// All values needs to be defined here in this way
	viraj := person{"Viraj", "Shelke", contactInfo{"vrjs.29@gmail.com", 421202}, optionalInfo{"Akshay"}}

	sid := person{
		firstName: "Sid",
		lastName:  "Dalvi",
		contact: contactInfo{
			email: "test@test.com",
			zip:   421001,
		},
	}

	/*
		If the structure is not initialized it will have zero values assigned depending on datatype
		So,
		string will have blank ""
		int will have 0
		float will have 0
		bool will have false
	*/
	// var tasmai person
	tasmai := person{}

	fmt.Println("Person Viraj", viraj)
	fmt.Println("Person Sid", sid)
	fmt.Println("Person Tasmai", tasmai)
	fmt.Printf("Person Tasmai %+v\n", tasmai)

	tasmai.firstName = "Tasmai"
	tasmai.lastName = "Shelke"
	tasmai.contact.email = "tasz@gmail.com"
	tasmai.contact.zip = 421202

	fmt.Printf("Person Tasmai updated %+v\n", tasmai)
	viraj.print()
	// The following function is in capable of updating the struct field value
	viraj.updateName("VIRAJJ")
	viraj.print() // This print prints the same as that of the above one
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName // somehow has no effect - refer pointers
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
