package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Variadic function allows to pass variable no of argruments to a function.
	// ... ie. 'ellipse' is used to specify variadic function.

	simpleVariadic()
	simpleVariadic("ashish")
	simpleVariadic("ashish", "arya")

	mixedVariadic(3, "ashish", "arya")

	complexVariadic(3, "ashish", []string{"ashish", "arya"}, true)


	// Defer in golang
	// - Allows us to schedule function call after the function completes.
	// - Used with open-close, connect-disconnect, lock-unlock.
	// - Defer function run even panic occurs.
	// - Defer executes in LIFI order.

	defer second()
	first()

	deferLifo()
	fmt.Println()

	fmt.Println("Thanks, see you in next tutorial.")
}

func simpleVariadic(s ...string) {
	fmt.Println(s)
}

func mixedVariadic(i int, s ...string) {
	fmt.Println("\nMixed variadic:")

	for index := range s {
		fmt.Print(s[index], " ")
	}

	fmt.Println()
}

func complexVariadic(i ...interface{}) {
	fmt.Println("\nPrinting Complex variadic:")
	for _, val := range i {
		fmt.Println(val, reflect.TypeOf(val))
	}
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func deferLifo() {
	for i:=0;i<5;i++ {
		defer fmt.Printf("%d ", i)
	}
}