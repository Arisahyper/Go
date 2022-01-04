package main

import "fmt"

func main() {
	list := []string{"a", "b", "c"}
	for i, v := range list {
		fmt.Printf("%d: %s\n", i, v)
	}
}
