package main

import "fmt"

func main() {
	var arr [10]int
	arr[0] = 10
	arr[1] = 20

	anotherArr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var yetAnotherArr [4]int = [4]int{1, 2, 3, 4} // This builds an array and then slice referencing to it (see below)

	fmt.Println(arr, arr[0], arr[1])
	fmt.Println(anotherArr, anotherArr[0], anotherArr[1])
	fmt.Println(yetAnotherArr, yetAnotherArr[0], yetAnotherArr[1])
	fmt.Println([10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, [10]int{0: 10, 1: 20}[0], [10]int{0: 10, 1: 20}[1])
	fmt.Println()

	// Array slices
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[1:4])
	fmt.Println(primes[1:])
	fmt.Println(primes[:4])

	// Warning: Array slices are references to underlying array!
	a := primes[1:4]
	a[0] = 10
	a[1] = 20
	fmt.Println(primes, len(a), cap(a)) // len() - length of slice, cap() - capacity of the whole array
	fmt.Println()

	// Dynamic allocation of slices
	allocated := make([]int, 10)
	allocated[0] = 10
	allocated[1] = 20
	fmt.Println(allocated)

	// Slices are higher level object based on underlying array. Underlying array can be reallocated to accommodate new
	// values. In this case, a new slice based on a new array will be returned. Otherwise (if the array is long enough
	// to accommodate new values) we will get a new slice based on the old array.
	allocated = append(allocated, 123, 234, 345)
	allocated = append(allocated, 123456)
	fmt.Println(allocated)

	// Slice of slices
	matrix := [][]int{
		[]int{10, 20, 30, 40},
		[]int{11, 21, 31, 41},
		[]int{12, 22, 32, 42},
		[]int{13, 23, 33, 43},
	}
	fmt.Println(matrix, matrix[0][0])
	fmt.Println()
}
