package main 

import (
	"fmt"
)

func main() {
	// Go panics occurs at out of bound or nil point dereference.
	// It panic the go routine and exits the program execution
	// Used to handle logically 'impossible' situations

	var input int
	fmt.Scan(&input)

	switch input {
	case 1:	
		fmt.Println("You've put 1.")
	case 2:
		fmt.Println("You've put 2.")
	default:
		panic(fmt.Sprintf("Wrong input %d!!!", input))
	}

	// recover() collects the panic error from panic
	defer func() {
		action := recover()
		fmt.Println(action)
	}()
}