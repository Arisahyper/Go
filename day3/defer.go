package main

import "fmt"

func deferWrapper() {
	defer fmt.Println("2")

	fmt.Println("1")
}

func main() {
	defer fmt.Println("4")

	deferWrapper()

	fmt.Println("3")
}

/*
そのブロック内（今回は関数）のdeferは、ブロックの最後に実行される。
*/
