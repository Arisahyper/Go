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

func main() {
	ans, _ := sqrt(2)
	fmt.Println(ans)
}
