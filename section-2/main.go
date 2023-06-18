package main

import "fmt"

// define the type
type person struct {
	firstName string
	lastName  string
}

func main() {
	// declare
	// var ehsan person

	// method-1

	// initialize
	// ehsan.firstName = "Ehsan"
	// ehsan.lastName = "Ullah"

	// method-2

	// initialize
	// ehsan = person{
	// 	firstName: "Ehsan",
	// 	lastName:  "Ullah",
	// }

	// declaration and initialization
	ehsan := person{
		firstName: "Ehsan",
		lastName:  "Ullah",
	}

	fmt.Println(ehsan)
}
