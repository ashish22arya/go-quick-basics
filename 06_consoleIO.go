// Go console input output

package main

import (
	"fmt"
)

func main() {
	// Printing to std.out
	i, s, f, b := 32, "Ashish", 43.32, true

	fmt.Print(i, s, f, b, "\n")			// writes to stdandard output using default format
	// space omited if nither is string
	
	fmt.Println(i, s, f, b)				// Writes to std output followed by new line.
	// space is added between operand

	fmt.Printf("%d %T\n", i, i)		// %b for base 2
	fmt.Printf("%s %T\n", s, s)		// %q for double quoted string
	fmt.Printf("%f %T\n", f, f)		// %e or %g for large exponent
	fmt.Printf("%t %T\n", b, b)

	// %v	the value in a default format
	fmt.Printf("%#v %T\n", f, f)

	// %#v	a Go-syntax representation of the value
	// %T	a Go-syntax representation of the type of the value
	
	type info struct {
		name 	string
		age 	int
	}
	in := info{name: "Ashish", age: 24}
	// when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("%+v %T\n", in, in)

	// Width and precision of floating point number
	fmt.Printf("%f\n", f)		// %f default width and %g precision
	fmt.Printf("%4f\n", f)		// with width 4
	fmt.Printf("%.1f\n", f)		// with precision 1
	fmt.Printf("%3.4f\n", f)	// with width 3 and precision 4

	// Read from standard input
	fmt.Scan(&i)			// Consider new line as spaces
	fmt.Scanln(&s, &b)		// Stops scanning at new line
	fmt.Scanf("%f", &f) 	// Formatted scan

	fmt.Println(i, f)

	// Fscan, Fscanf, Fscanln -> Scans from io reader.
	// Sscan, Sscanf, Sscanln -> Scans from argument string
	
	fmt.Println("Thank you, see you in next tutorial.")
}