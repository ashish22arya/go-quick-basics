package main

import (
	"fmt"
)

func main() {
	// Array is used to store values of same type.
	// It is of fixed length and cannot increase its size as required.

	// 1. Array declaration
	var arr [5]int
	arr[0] = 2		// Assigning a value
	fmt.Println("arr :", arr)

	// 2. Array declaration and define at same time
	arr1 := [5]int{1, 2, 3, 4, 5}		
	fmt.Println("arr1 :", arr1)

	var arr2 [5]int = [5]int{1, 2, 3}
	fmt.Println("arr2 :", arr2)

	// 3. Array definition omiting the size. Size is determined by no of values.
	arr3 := [...]int{1, 2, 3}
	fmt.Println("arr3 :", arr3)

	// 4. Initialize values for specific indexes
	arr4 := [5]int{1:2, 4:5}
	fmt.Println("arr4 :", arr4)

	// 5. Iteration over an array
	fmt.Println("Iteration over array.")
	for i := range arr4 {
		fmt.Print(arr4[i], " ")
	}
	fmt.Println()

	// Iteration: for index, val := range arr4 {}
	// Iteration: for _, val := range arr4 {}
	// Iteration: for range arr4 {}

	// Passing array by value and reference
	arr5 := arr4
	arr6 := &arr4
	arr4[0] = 5

	fmt.Println("arr4 :", arr4)
	fmt.Println("arr5 :", arr5)
	fmt.Println("arr6 :", arr6)

	// Partial print of array
	fmt.Println("arr4[:]", arr4[:])
	fmt.Println("arr4[:3]", arr4[:3])
	fmt.Println("arr4[2:]", arr4[2:])
	fmt.Println("arr4[1:4]", arr4[1:4])

	fmt.Println("Thanks, see you in next tutorial.")
}