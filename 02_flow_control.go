package main

import "fmt"
import "runtime"

func main() {
	// For
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	// While
	sum := 0
	for sum < 10 {
		fmt.Println(sum)
		sum += 1
	}
	fmt.Println()

	// While truen and If
	sum = 0
	for {
		fmt.Println(sum)
		sum += 1
		if sum == 10 {
			break
		}
	}
	fmt.Println()

	// If with a short statement
	if a := 10; a < 20 {
		fmt.Println(a)
	}
	fmt.Println()

	// Switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Good for you!")
	case "linux":
		fmt.Println("Not so bad.")
	default:
		fmt.Println(":-(")
	}
	fmt.Println()

	// Switch with no condition
	a, b, c := 10, 20, 30
	switch {
	case a < b:
		fmt.Println(":-)")
	case b < c:
		fmt.Println(":-(")
	} // Switch will choice only one case
	fmt.Println()

	// Defer
	deferExample()
}

func deferExample() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}
