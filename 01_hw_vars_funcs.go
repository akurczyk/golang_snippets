package main

import "fmt"

var a, b, c, d = 1, 2, 3, 4
var x, y, z string

func main() {
	// Display Hello World!
	fmt.Println("Hello World!")
	fmt.Println()

	// Print global variables declared without type declaration
	fmt.Println(a, b, c, d)
	fmt.Println()

	// Assign values to global strings and print them.
	// Empty variables after declaration are 0, false or "".
	x = "abc"
	y = "def"

	fmt.Println(x, y, z)

	// Declare local variables and print them
	var i int = 10
	var f float64 = 20.5
	var u uint = 30
	var s string = "Lorem ipsum"

	fmt.Println(i, f, u, s)
	fmt.Println()

	// Print local previously declared local variables along with type definition using fmt.Printf
	fmt.Printf("%T: %v\n%T: %v\n%T: %v\n%T: %v\n\n", i, i, f, f, u, u, s, s)

	// Create new variables with short assigment, then print them
	new_i := i
	new_f := f

	fmt.Println(new_i, new_f)
	fmt.Println()

	// Convert types
	var tmp_a int = 10
	var tmp_b float64 = 3.1415
	var tmp_c float64 = float64(tmp_a)
	var tmp_d int = int(tmp_b)

	fmt.Println(tmp_a, tmp_b, tmp_c, tmp_d)
	fmt.Println()

	// Call functions
	fmt.Println("another_func(10, 10):", another_func(10, 10))
	fmt.Println("yet_another_func(10, 10):", yet_another_func(10, 10))
}

func another_func(x, y int) int {
	return x + y
}

func yet_another_func(x, y int) (s int) {
	s = x + y
	return
}

/*
   Go types:

   bool

   string

   int  int8  int16  int32  int64
   uint uint8 uint16 uint32 uint64 uintptr

   byte // alias for uint8

   rune // alias for int32
        // represents a Unicode code point

   float32 float64

   complex64 complex128
 */
