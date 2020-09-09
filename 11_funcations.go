package main

import "fmt"

func main() {
	// Function is used to extract commonly used blocks.
	simpleFunc()

	// Function with argument
	findArea(3, 5)

	// Function with return
	area := findTotalArea(3, 6)
	fmt.Println("Total area is :", area)

	// Function with multiple return
	area, peri := findAreaPerimeter(3, 5)
	fmt.Println("Area :", area, "Perimeter :", peri)

	// Passing address to the function
	name, age := "Ashish", 24
	updateData(&name, &age)
	fmt.Println("Updated data: ", name, age)

	// Anonymous function : Function with no identifier
	func(i int) {
		fmt.Println("Number :", i)
	}(5)

	// Closures function: Anonymous function that access outside variables
	func() {
		fmt.Println("User: ", name, age)
	}()

	for i:=0;i<3;i++ {
		multiple := func() int {
			return i * 10
		}()
		fmt.Println("Multiple :",multiple)
	}

	// High order functions: Function that accept function as an argument.


	
}

func simpleFunc() {
	fmt.Println("This is simplest function.")
}

func findArea(length int, bredth int) {
	fmt.Println("Area is :", length * bredth)
}

func findTotalArea(length, breadth int) int {
	return length * breadth
}

func findAreaPerimeter(length, breadth int) (area, peri int) { // Named return value
	area = length * breadth
	peri = 2 * (length + breadth)
	return 
}

func updateData(name *string, age *int) {
	*name = *name + " Arya"
	*age += 2
	return
}