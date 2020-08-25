package main

import (
	"errors"
	"fmt"
)

func double(n int) (int, error) {
	if n >= 0 {
		return n * 2, nil
	} else {
		return -1, errors.New("you can double only positive numbers")
	}
}

func main() {
	var value int
	var err error

	value, err = double(10)
	fmt.Println(value, err)

	value, err = double(-10)
	fmt.Println(value, err)

	if err != nil {
		panic(err)
	}

	fmt.Println("This part of the code won't be executed.")
}
