package main

import "fmt"

func main() {
	// For loop: Loop to iterate several times or over slice, map, string, array

	// Different syles 
	for i := 0; i < 5 ; i++ {
		fmt.Println("Hello")
	}

	i := 0
	for ; i<5; i++ {		// Missing initializer
		fmt.Println("Hii")
	} 

	for i := 0; ; i++ {		// Missing condition check
		fmt.Println("Namaste")
		if i == 5 {
			break
		}
	}

	for i := 0; i<5; {		// Missing incrementor
		fmt.Println("Hi")
		i++
	}

	// For loop over string
	name := "Ashish"
	for i := range name {
		fmt.Println("Hi", string(name[i]))	// name[i] gives the ASCII for each character
	}

	// Infinite loop
	j := 0  
	for {
		fmt.Println("Thanks")
		if j == 2 {
			break
		}
		j++
	}

	// For loop: Range
	grade := map[string]string{"A":"Excellent", "B":"Good", "C":"Average"}
	for key := range grade {
		fmt.Println(key)
	}

	for key, val := range grade {
		fmt.Println(key, val)
	}

	fmt.Println("Thanks, see you in next tutorial.")
}