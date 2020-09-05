// Go Data types

package main 

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// Basic data types: int, string, float(by default float64) and bool

	var i, s, f, b = 5, "Data types", 4.0, true

	fmt.Println("Type :",reflect.TypeOf(i), reflect.TypeOf(s), reflect.TypeOf(f), reflect.TypeOf(b))
	fmt.Println("Bytes: ",unsafe.Sizeof(i), unsafe.Sizeof(s), unsafe.Sizeof(f), unsafe.Sizeof(b))
	// Memory efficient data types as per requirement.
	// Integer type
	var (
		i0 int = 0
		i1 uint8 = 1
		i2 uint16 = 2
		i3 uint32 = 3
		i4 uint64 = 4	
		i5 int8 = 5
		i6 int16 = 6
		i7 int32 = 7
		i8 int64 = 8
	)

	fmt.Println("int	Range: -2^64 - 2^64-1		Bytes", unsafe.Sizeof(i0))
	fmt.Println("uint8	Range: 0 - 255			Bytes", unsafe.Sizeof(i1))
	fmt.Println("uint16 	Range: 0 - 65535		Bytes", unsafe.Sizeof(i2))
	fmt.Println("uint32 	Range: 0 - 4294967295		Bytes", unsafe.Sizeof(i3))
	fmt.Println("uint64 	Range: 0 - 18446744073709551615	Bytes", unsafe.Sizeof(i4))
	fmt.Println("int8	Range: -2^8 - 2^7-1		Bytes", unsafe.Sizeof(i5))
	fmt.Println("int16 	Range: -2^16 - 2^16-1		Bytes", unsafe.Sizeof(i6))
	fmt.Println("int32 	Range: -2^32 - 2^32-1		Bytes", unsafe.Sizeof(i7))
	fmt.Println("int64 	Range: -2^64 - 2^64-1		Bytes", unsafe.Sizeof(i8))

	// Floating type
	var (
		f1 float32 = 4.6
		f2 float64 = 4.8
	)
	fmt.Println("float32 	Bytes", unsafe.Sizeof(f1))
	fmt.Println("float64 	Bytes", unsafe.Sizeof(f2))

	// Complex type
	var (
		c1 complex64 = 4 + 5i   	// Complex no with float32 real and imaginary part
		c2 complex128 = 5 + 6i		// Complex with float64 real and imaginary part
	)
	fmt.Println("complex64 	Bytes", unsafe.Sizeof(c1))
	fmt.Println("complex128 	Bytes", unsafe.Sizeof(c2))

	fmt.Println("Thanks, see you in next tutorial.")

}