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

// 倍する
func (v *Vertex) Scale(i int) {
	v.X *= i
	v.Y *= i
}

func main() {
	v := Vertex{3, 4}

	fmt.Println(Area(v))
	fmt.Println(v.Area())

	fmt.Println(v)
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Area())
}
