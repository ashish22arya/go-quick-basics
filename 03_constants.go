// Go constants

package main

import (
	"fmt"
)

func main() {
	// Constant declaration
	const CITY string = "Indore"
	const CODE = 91

	fmt.Println(CITY, CODE)

	// Constant block
	const (
		AREA = "Vijay Nagar"
		PINCODE = 452010
	)

	fmt.Println(AREA, PINCODE)

	/******** More on constants *********/
	// - Declaration is using valid identifier.
	// - Must be initialized at time of declaration, as can not change value later.
	// - By convention, use all capital letters for quick identificaion in source code.

	fmt.Println("Thanks, see you in next tutorial.")
}