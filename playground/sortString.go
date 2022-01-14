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
}
