package main

import "fmt"

func main() {
	// Function in variable
	myFunc := func(x, y int) int {
		return x + y
	}
	fmt.Println(myFunc(10, 20))

	// Function in argument
	myFunc2 := func(x, y int, fn func(int, int) int) int {
		return fn(x, y)
	}
	fmt.Println(myFunc2(10, 20, myFunc))

	// Function as return value
	myFunc3 := func(a int, operation string) func(int, int) int {
		switch operation {
		case "add":
			return func(x, y int) int {return x + y + a}
		case "sub":
			return func(x, y int) int {return x - y + a}
		case "mul":
			return func(x, y int) int {return x * y + a}
		case "div":
			return func(x, y int) int {return x / y + a}
		case "mod":
			return func(x, y int) int {return x % y + a}
		default:
			return func(int, int) int {return a}
		}
	}

	adder := myFunc3(10, "add")
	subtractor := myFunc3(10, "sub")

	fmt.Println(adder(10, 10))
	fmt.Println(subtractor(10, 10))

	fmt.Println(myFunc3(10, "mul")(10, 10))
	fmt.Println(myFunc3(10, "div")(10, 10))
	fmt.Println(myFunc3(10, "mod")(5, 10))
	fmt.Println()
}
