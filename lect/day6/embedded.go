package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// 継承
type Vertex3D struct {
	Vertex
	Z int
}

func (v Vertex3D) Area3D() int {
	return v.X * v.Y * v.Z
}

func (v *Vertex3D) Scale3D(i int) {
	v.X *= i
	v.Y *= i
	v.Z *= i
}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	v := New(3, 4, 5)

	fmt.Println(v)
	fmt.Println(v.Area3D())

	v.Scale3D(10)
	fmt.Println(v)
	fmt.Println(v.Area3D())
}
