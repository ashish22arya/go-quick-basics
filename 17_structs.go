package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Struct is used to define user defined data types.
	/*
	type identifier struct {
		field1 	datatype
		field2 	datatype
	}
	*/

	type city struct {
		name 	string
		pincode	int
	}

	var c1 city			// Instance of struct
	c1.name = "Indore"
	c1.pincode = 452010

	fmt.Println("Struct c1:", c1, " -> ", reflect.ValueOf(c1).Kind())

	// Instance of struct using struct literal
	var c2 = city{"Shajapur", 465001}
	fmt.Println("Struct c2:", c2)

	// Initialize using new keyword
	var c3 = new(city)
	c3.name = "Berchha"
	c3.pincode = 465110
	fmt.Println("Struct c3:", c3)

	// Pointer to struct
	var c4 = &city{"Dewas", 2353453}
	fmt.Println("Struct c4: ", c4)

	// Nested struct
	type state struct {
		name 	string
		cities	[]city
	}
	
	d1 := state {
		name : "Madhya Pradesh",
		cities : []city{c1, c2},
	}

	fmt.Println("District: ", d1)

	// Struct copy to new structure
	// Copy by value : var c5 = c1
	// Copy by reference : var c5 = &c1

	// Adding method to struct type
	// func(c city) printCity() {}
	// func(c *city) editCity() {}

	fmt.Println("Thanks, see you in next tutorial.")
}