package main

import "fmt"

func even(i int) bool {
	return i%2 == 0
}

func main() {
	for i := 0; i < 100; i++ {
		if even(i) {
			fmt.Println(i)
		}
	}
}

// 偶数の絞り込み
