package main

import "fmt"

type Vertex struct {
	x int
	y int
}

// classのnewみたいなもの
func New(x int, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	v := New(3, 4)
	fmt.Println(v)
}
