package main

import "fmt"

func main() {
	// Decision making: if, if .. else, if .. else if ... else, switch
	// if conditioning
	name := "Ashish"
	if name == "Ashish" {
		fmt.Println("Correct name.")
	}

	// if .. else conditioning
	if name == "Arya" {
		fmt.Println("Correct name.")
	} else {
		fmt.Println("Incorrect name.")
	}

	// if .. else if .. else conditioning
	name = "Arya"
	if name == "Ashish" {
		fmt.Println("It's a name.")
	} else if name == "Arya" {
		fmt.Println("It's a surname")
	} else {
		fmt.Println("Incorrect name.")
	}

	// if supports initialization of varible precceding condition
	if city := "Indore"; city == "Indore" {
		fmt.Println("City:", city)
	}

	fmt.Println("Thanks, see you in next tutorial.")
}