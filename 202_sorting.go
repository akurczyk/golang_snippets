package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	// https://gobyexample.com/sorting

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	fmt.Println(sort.StringsAreSorted(strs))

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)
	fmt.Println(sort.IntsAreSorted(ints))

	// https://gobyexample.com/sorting-by-functions

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println("By function:", fruits)
	fmt.Println(sort.IsSorted(byLength(fruits)))
}
