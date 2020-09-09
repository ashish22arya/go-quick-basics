package main

import "fmt"

func main() {
	// Arithmetic operators: +, -, *, /, %, ++, -- 
	fmt.Println("\nArithmetic operators: +, -, *, /, %, ++, --")
	i, j := 5, 2
	fmt.Println("i + j :", i+j)
	fmt.Println("i - j :", i-j)
	fmt.Println("i * j :", i*j)
	fmt.Println("i / j :", i/j)
	fmt.Println("i % j :", i%j)
	i++; j--				// ; is used to terminate the expression 
	fmt.Println("i++ :", i)
	fmt.Println("j-- :", j)

	// Assignment operators: =, +=, -=, *=, /=, %=
	fmt.Println("\nAssignment operators: =, +=, -=, *=, /=, %=")
	i = 43; j = 5
	i += j
	fmt.Println("i += j :", i)
	
	i = 43; j = 5
	i -= j
	fmt.Println("i -= j :", i)

	i = 43; j = 5
	i *= j
	fmt.Println("i *= j :", i)

	i = 43; j = 5
	i /= j
	fmt.Println("i += j :", i)

	i = 43; j = 5
	i %= j
	fmt.Println("i %= j :", i)

	// Comparision operators: ==, !=, <, >, <=, >=
	fmt.Println("\nComparision operators: ==, !=, <, >, <=, >=")

	i = 43; j = 8
	fmt.Println("i == j :", i == j)
	fmt.Println("i != j :", i != j)
	fmt.Println("i < j :", i < j)
	fmt.Println("i > j :", i > j)
	fmt.Println("i <= j :", i <= j)
	fmt.Println("i >= j :", i >= j)

	// Logical operators: &&, ||, !
	fmt.Println("\nLogical operators: &&, ||, !")
	i = 5; j = 3
	fmt.Println("Logical AND && :", (i > j && j == 3))
	fmt.Println("Logical OR || :", (i < j && i == 3))
	fmt.Println("Logical NOT ! :", !(i < j))

	// Bitwise operators: &, |, ^, <<, >>
	fmt.Println("\nBitwise operators: &, |, ^, <<, >>")
	var x uint = 9  //0000 1001
	var y uint = 65 //0100 0001
	var z uint

	z = x & y
	fmt.Println("x & y  =", z)

	z = x | y
	fmt.Println("x | y  =", z)

	z = x ^ y
	fmt.Println("x ^ y  =", z)

	z = x << 1
	fmt.Println("x << 1 =", z)

	z = x >> 1
	fmt.Println("x >> 1 =", z)

	fmt.Println("Thanks, see you in next tutorial.")
}