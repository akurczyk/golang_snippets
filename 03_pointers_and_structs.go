package main

import "fmt"

func main() {
	// Pointers
	var i, j int = 10, 20
	var k, l *int = &i, &j

	fmt.Println(i, j, *k, *l) // 10 20 10 20
	fmt.Println(&i, &j, k, l) // <ptr> <ptr> <ptr> <ptr>
	fmt.Println()

	// Structs
	type MyStruct struct {
		A int
		B float64
		C string
	}

	var myInstance MyStruct
	myInstance.A = 10
	myInstance.B = 3.1415
	myInstance.C = "Lorem ipsum"

	var mySecondInstance MyStruct = MyStruct{20, 10.5, "Aaa"}

	fmt.Println(myInstance)
	fmt.Println(mySecondInstance)
	fmt.Println(MyStruct{30, 12.3, "Bbb"})
	fmt.Println(MyStruct{A: 40}) // B=0.0, C=""
	fmt.Println()
}
