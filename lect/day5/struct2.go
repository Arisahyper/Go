package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
	// (*v).X = 1000 / こういう書き方もある
}

func main() {
	v := Vertex{1, 2}
	changeVertex(v)
	fmt.Println(v) // {1 2}	/ 値渡し

	changeVertex2(&v)
	fmt.Println(v) // {1000 2}	/ 参照渡し
}
