package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits []string
	fruits = []string{"banana", "cherry", "apple", "orange", "grape"}
	fmt.Println(fruits) // normal

	sort.Strings(fruits) // sort string
	fmt.Println(fruits)  // sorted

	sort.Sort(sort.Reverse(sort.StringSlice(fruits))) // reverse string
	fmt.Println(fruits)                               // reverse

	fruits = reverse(fruits) // reverse string
	fmt.Println(fruits)      // reverse
}

// 非破壊的なソート
func reverse(s []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(s))) // reverse string
	return s
}