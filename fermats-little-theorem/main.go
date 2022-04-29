package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println("a =",i ,f(float64(i), 7))
	}
}

func f(a float64, p float64) bool {
	// (mod p)
	return math.Mod(math.Pow(a, (p-1)), p) == 1
}

// 1=1≡1(mod7)16=1≡1(mod7)
// 2=64≡1(mod7)26=64≡1(mod7)
// 3=729≡1(mod7)36=729≡1(mod7)
// 4≡(−3)6=36≡1(mod7)46≡(−3)6=36≡1(mod7)
// 5≡(−2)6=26≡1(mod7)56≡(−2)6=26≡1(mod7)
// 6≡(−1)6=16≡1(mod7)66≡(−1)6=16≡1(mod7)
