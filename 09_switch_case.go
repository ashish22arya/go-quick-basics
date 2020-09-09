package main

import "fmt"

func main() {
	// Switch is an alternative to multiple if .. else statement

	name := "Ashish"
	switch name {
		case "Ashish":
			fmt.Println("It's a name.")
		case "Arya":
			fmt.Println("It's a surname.")
		default:
			fmt.Println("Not a vaild name!")
	}

	// Switch with fallthrough : Force program execution to next case.
	fmt.Println("Fallthrough example:")
	switch name {
		case "Ashish":
			fmt.Println("It's a name.")
			fallthrough
		case "Arya":
			fmt.Println("It's a surname.")
		default:
			fmt.Println("Not a vaild name!")
	}

	// Switch with multiple cases
	// ex: 	case "Ashish", "ashish":
	// 		case 4,5:

	// Switch with condition statements
	// ex. 	case value < 10:
	//		case len(name) < 13:
		
	// Swithc with initializer: variable can be accesses inside switch
	// ex. 	switch name:="Ashish"; name {
		
	fmt.Println("Thanks, see you in next tutorial.")
}