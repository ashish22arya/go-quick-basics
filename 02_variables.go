// Golang variables

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Varible declaration and initialization
	var name string
	var age int

	name = "Ashish"
	age = 24
	
	fmt.Println(name, age)

	// Declaration and initialization at same time
	var name1 string = "Ashish"
	var age1 int = 24

	fmt.Println(name1, age1)

	// Declaration and initialization omitting data type
	var name2 = "Ashish"
	var age2 = 24

	fmt.Println(name2, age2)

	// Short hand for declaration and initialization
	name3 := "Ashish"		// Type of variable is as same as type of value assigned
	age3 := 24		

	fmt.Println(name3, reflect.TypeOf(name3))
	fmt.Println(age3, reflect.TypeOf(age3))

	// Declare multiple variables
	var add1, add2 string = "Indore", "MP"
	pincode, country := 453234, "India"

	fmt.Println(add1, add2, pincode, country)

	// Variable blocks (for better readability)
	var (
		name4 = "Ashish"
		age4 = 24
	)

	fmt.Println(name4, age4)

	/**************** More information *************/
	// - Variable name must be a valid identifier(start with letter and can have multiple letters/digits)
	// - Scope is defined into the block only.
	// - Variable declared inside main() must be use.
	// - Variable name start with Capital letter can be exported outside of package.
	// - Variable name start with small letter must be used inside package.
	// - Variable names are case sensitive.

	fmt.Println("Thanks, see you in next tutorial.")
}