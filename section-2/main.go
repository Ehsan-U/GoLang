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
