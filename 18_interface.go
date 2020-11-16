package main

import (
	"fmt"
)

type user interface {
	printName()
	showUserNumber(i int) int
}

type manager string

func (m manager) printName() {
	fmt.Println("Manager")
}

func (m manager) showUserNumber(i int) int {
	return i
}

func main() {
	// Interface is an abstract type that describes all the method in method set.
	/*
	type identifier interface {
		method1()
		method2()
	}
	*/

	var m1 manager
	m1 = manager("Ashish")

	m1.printName()
	fmt.Println("Number:", m1.showUserNumber(5))

	// Additionals on interface
	// 1. Using type assertion, we can use multiple interfaces.
	// 2. Multiple interfaces can defince same methods in list of method set.
	// 3. Interface allows to have interfaces accepting address of variables.
	// 4. If a method accepts a type value, then the interface must receive a type value.
	// 5. If a method has a pointer receiver, then the interface must receive the address 
	// of the variable of the respective type.
	// 6. interface{} - empty interface, and it is used to accept values of any type.
	// 7. Polymorphism is achieved using interfaces.

	fmt.Println("Thanks, see you in next tutorial.")
}