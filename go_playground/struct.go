// Golang program to show how to
// declare and define the struct

package main

import "fmt"

// Defining a struct type
type Address struct {
	Name    string
	city    string
	Zipcode int
}

func main() {

	// Declaring a variable of a `struct` type
	// All the struct fields are initialized
	// with their zero value
	var a0 = Address{"Muang", "Surathani", 84000}

	fmt.Println(a0)

	// Declaring and initializing a
	// struct using a struct literal
	a1 := Address{"Phayathai", "BKK", 10400}

	fmt.Println("Address1: ", a1)

	// Naming fields while
	// initializing a struct
	a2 := Address{Name: "ladprao", city: "BKK", Zipcode: 10230}

	fmt.Println("Address2: ", a2)

	fmt.Println(a0.Name)
	// Uninitialized fields are set to
	// their corresponding zero-value

	// Can't Mixture of struct like this
	/*
			a3 := Address{Name: "Dindang", "BKK", Zipcode: 10400}

		    	fmt.Println("Address3: ", a3)
	*/
}
