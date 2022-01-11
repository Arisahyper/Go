package main

import (
	"fmt"
	"math"
)

func sliceScan() [][2]int {
	var n int
	fmt.Scan(&n)

	var list [][2]int
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)

		var vertex [2]int
		vertex[0] = x
		vertex[1] = y
		list = append(list, vertex)
	}
	return list
}

// 線分の長さを求める関数
func length(x1, y1, x2, y2 int) int {
	return int(math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))))
}

func main() {
	list := sliceScan()

	fmt.Println(list)
	fmt.Println(length(0, 0, 1, 1))
}
