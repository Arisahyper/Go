package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

func main() {
	v := Vertex{X: 1, Y: 2} // 初期化
	fmt.Println(v)          // {1 2}
	fmt.Println(v.X, v.Y)   // 1 2
	v.X = 100               // 代入
	fmt.Println(v.X, v.Y)   // 100 2

	v2 := Vertex{X: 1} // Y:0
	fmt.Println(v2)    // {1 0}

	v3 := Vertex{1, 2, "test"} // 省略可能
	fmt.Println(v3)            // {1 2 test}

	v4 := Vertex{}
	fmt.Println(v4) // {0 0  }
}
