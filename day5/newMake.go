package main

import "fmt"

func main() {
	s := make([]int, 0)
	fmt.Printf("%T\n", s)	// []int

	m := make(map[string]int)
	fmt.Printf("%T\n", m)	// map[string]int
}
