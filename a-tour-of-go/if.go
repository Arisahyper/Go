package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("x must be >= 0, but got %v", x)
	}
	return math.Sqrt(x), nil
}

func pow(x, n float64) float64 {
	return math.Pow(x, n)
}

func main() {
	ans, _ := sqrt(2)
	fmt.Println(ans)

	ans = pow(2, 3)
	fmt.Println(ans)
}
