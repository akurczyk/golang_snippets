package main

import "fmt"

func main() {
	var myMap map[string]string = make(map[string]string) // map[<key type>]<value type>
	// or: var myMap = make(map[string]string)
	// or: myMap := make(map[string]string) (this can be done only in functions)
	myMap["abc"] = "123"
	myMap["def"] = "456"
	fmt.Println(myMap)
	fmt.Println()

	anotherMap := map[string]int{
		"ABC": 123,
		"CDE": 345,
		"EFG": 567,
	}
	fmt.Println(anotherMap)
	fmt.Println()

	type Coords struct {
		Lat float64
		Lon float64
	}
	yetAnotherMap := map[string]Coords{
		"Budapest": {47.497913, 19.040236}, // or: Coords{...}
		"Praha": {50.073658, 14.418540},
	}
	fmt.Println(yetAnotherMap)
	fmt.Println()

	delete(yetAnotherMap, "Budapest")
	value1, isPresent1 := yetAnotherMap["Budapest"]
	value2, isPresent2 := yetAnotherMap["Praha"]
	fmt.Println(value1, isPresent1)
	fmt.Println(value2, isPresent2)
	fmt.Println()
}
