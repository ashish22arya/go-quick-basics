package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Slices are dynamic array that can increase and decrease its size.
	// Can be accessed using index same as array

	// 1. Define an empty slice
	var slice []int
	fmt.Println("slice :", reflect.TypeOf(slice).Kind())

	// 2. Make slice using make function
	var slice2 = make([]int, 3)		// Here capacity and length is same 
	fmt.Println("slice2 : len:", len(slice2), "cap:", cap(slice2))

	var slice3 = make([]int, 3, 5)  // Slice of lenght 3 and capacity 5
	fmt.Println("slice3 : len:", len(slice3), "cap:", cap(slice3))

	// 3. Initialze slice with values
	var slice4 = []int{1, 3, 5}
	fmt.Println("slice4 :", slice4, "len:", len(slice4), "cap:", cap(slice4))

	// 4. Create slice using new keyword
	var slice5 = new([5]int)
	fmt.Println("slice5 :", slice5, "len:", len(slice5), "cap:", cap(slice5))

	// 5. Adding an item to slice using append function
	slice4 = append(slice4, 7)
	slice4 = append(slice4, 9, 11, 13)
	fmt.Println("slice4 :", slice4, "len:", len(slice4), "cap:", cap(slice4))

	// 6. Slicing element of a slice
	// Ex. slice4[:], slice4[1:3], slice4[3:]

	// 7. Append two slices into one
	slice2 = append(slice2, slice4...)
	fmt.Println("slice2 :", slice2, "len:", len(slice2), "cap:", cap(slice2))

	fmt.Println("Thanks, see you in next tutorial.")
}