package main

import (
	"fmt"
)

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	min := 1<<63 - 1

	for _, value := range l {
		if value < min {
			min = value
		}
	}
	fmt.Print(min)
}
