package main

import (
	"fmt"
	"sort"
)

func main() {
	// Maps are the collection of unordered pair of key-values.
	// Used to look up values by keys.

	// 1. Empty map
	var m = map[string]int{}
	m["Ashish"] = 1
	fmt.Printf("Map m : %v %T\n", m, m)

	// 2. Initialize map at declaration
	var m2 = map[string]int{"Arya": 2, "Goldy":3}
	fmt.Printf("Map m2 : %v\n", m2)

	// 3. Map declaration using make
	var m3 = make(map[string]int)
	m3["Indore"] = 4
	fmt.Printf("Map m3 : %v\n", m3)

	// 4. Map length
	fmt.Println("Map m3 length : ", len(m3))

	// 5. Accessing map items
	fmt.Println("Value associated with m3['Indore'] : ", m3["Indore"])

	// 6. Delete map items
	delete(m3, "Indore")
	fmt.Println("Map m3 length after delete element : ", len(m3))

	// 7. Iterate over map
	// Ex. for key := range m3 {}
	// Ex. for key, val := range m3 {}
	// Ex. for range m3 {}

	// 8. Truncate map
	for key := range m2 {
		delete(m2, key)
	}

	m = make(map[string]int)
	fmt.Println("Map m and m2 length after delete: ", len(m), len(m2))

	// Print Sorted map by key
	m["Bob"] = 2
	m["Ashish"] = 1

	var sortedKey = make([]string, 0, len(m))
	for key := range m {
		sortedKey = append(sortedKey, key)
	}
	sort.Strings(sortedKey)

	fmt.Println("Sorted by key.")
	for _, val := range sortedKey {
		fmt.Println(val, m[val])
	}

	fmt.Println("Thanks, see you in next tutorial.")
}