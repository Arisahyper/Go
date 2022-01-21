package main

import (
	"fmt"
	"project1/lib"
)

func main() {
	i := []int{1, 2, 3, 4, 5}
	result := lib.Average(i)
	fmt.Println(result)
}
