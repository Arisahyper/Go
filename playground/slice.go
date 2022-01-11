package main

import "fmt"

func main() {
	var h []string
	h = append(h, "hoge")
	fmt.Println(h)

	s := []string{"hoge", "fuga", "piyo"}
	fmt.Println(s)
	fmt.Println(len(s))	// length
	fmt.Println(cap(s)) // capacity
}
