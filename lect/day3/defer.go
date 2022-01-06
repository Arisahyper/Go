package main

import "fmt"

func deferWrapper() {
	defer fmt.Println("2")

	fmt.Println("1")
}

func stackingDefer() {
	fmt.Println("start")
	defer fmt.Println("3")
	defer fmt.Println("2")
	fmt.Println("1")
	fmt.Println("end")
}

func main() {
	stackingDefer()

	defer fmt.Println("4")

	deferWrapper()

	fmt.Println("3")
}

/*
そのブロック内（今回は関数）のdeferは、ブロックの最後に実行される。
*/
