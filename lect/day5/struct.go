package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{X: 1, Y: 2} // 初期化
	fmt.Println(v)          // {1 2}
	fmt.Println(v.X, v.Y)   // 1 2
	v.X = 100               // 代入
	fmt.Println(v.X, v.Y)   // 100 2
}
