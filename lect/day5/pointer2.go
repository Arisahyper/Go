package main

import "fmt"

// デリファレンス
func castOne(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	castOne(&n)
	fmt.Println(n)
}
