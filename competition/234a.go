package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return math.Pow(x, 2) + 2*x + 3
}

func main() {
	var n float64
	fmt.Scan(&n)

	fmt.Println(int(f(f(f(n)+n) + f(f(n)))))
}

// 関数を作成 -> 繰り返し処理
