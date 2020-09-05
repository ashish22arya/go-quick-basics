// Go type conversion

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1. String to Int
	strInt := "-2342"
	intVar, _ := strconv.Atoi(strInt)		// Atoi is same as ParseInt to int 
	intVar64, _ := strconv.ParseInt("0101010", 10, 64)	// Always return int64

	fmt.Printf("%v %T\n",intVar, intVar)
	fmt.Printf("%v %T\n",intVar64, intVar64)

	// 2. String to Float
	strFloat := "23.56"
	f64, _ := strconv.ParseFloat(strFloat, 64)		// Always return float64

	fmt.Printf("%v %T\n", f64, f64)

	// 3. String to boolean
	strBool := "True"
	b, _ := strconv.ParseBool(strBool)		
	// ParseBool accepts: 1,t,T,true,True,TRUE,0,f,F,false,False,FALSE
	fmt.Printf("%v %T\n", b, b)

	// 4. Bool to string
	b = false
	strBool = strconv.FormatBool(b)

	fmt.Printf("%v %T\n", strBool, strBool)

	// 5. Float to string
	f64 = 2342.234234234
	strFloat = strconv.FormatFloat(f64, 'f', -1, 64)

	fmt.Printf("%v %T\n", strFloat,strFloat)

	// 6. Int to string
	intVar = 52325
	strInt = strconv.Itoa(intVar)
	strInt2 := strconv.FormatInt(int64(intVar), 10)

	fmt.Printf("%v %T\n", strInt, strInt)
	fmt.Printf("%v %T\n", strInt2, strInt2)

	// 7. Int to int8, int 16, int32, int64
	var i int = 5
	i8 := int8(i)
	i16 := int16(i)
	i32 := int32(i)
	i64 := int64(i)

	fmt.Println("Converting int to int8/16/32/64")
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", i8, i8)
	fmt.Printf("%v %T\n", i16, i16)
	fmt.Printf("%v %T\n", i32, i32)
	fmt.Printf("%v %T\n", i64, i64)

	i32 = int32(i64)
	i16 = int16(i64)
	i8 = int8(i64)
	i = int(i64)

	fmt.Println("Converting int8/16/32/64 to int")
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", i8, i8)
	fmt.Printf("%v %T\n", i16, i16)
	fmt.Printf("%v %T\n", i32, i32)
	fmt.Printf("%v %T\n", i64, i64)

	// 8. Converting float32 <-> float64

	var f32 float32 = 234.24
	f64 = float64(f32)
	f32 = float32(f64)

	fmt.Printf("%v %T\n", f32, f32)
	fmt.Printf("%v %T\n", f64, f64)

	// 9. Converting int <-> float
	f32 = 242.243
	i32 = int32(f32)
	f64 = float64(i32)
	
	fmt.Printf("%v %T\n", i32, i32)
	fmt.Printf("%v %T\n", f32, f32)
	fmt.Printf("%v %T\n", f64, f64)
	
	f64 = float64(intVar)
	fmt.Printf("%v %T\n", intVar, intVar)
	fmt.Printf("%v %T\n", f64, f64)

	fmt.Println("Thanks, see you in next tutorial.")
}