package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// メソッド / Vertexに紐づく関数
func (v Vertex) Area() int {
	return v.X * v.Y
}

// 関数
func Area(v Vertex) int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}

	fmt.Println(Area(v))
	fmt.Println(v.Area())
}
