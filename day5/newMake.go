package main

import "fmt"

func main() {
	s := make([]int, 0)
	fmt.Printf("%T\n", s) // []int / 空のslice

	m := make(map[string]int)
	fmt.Printf("%T\n", m) // map[string]int / 空のmap

	var p *int = new(int)
	fmt.Printf("%T\n", p) // *int	/ ポインタ型
}
