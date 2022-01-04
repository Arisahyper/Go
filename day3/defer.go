package main

import "fmt"

func deferWrapper() {
	defer fmt.Println("2")

	fmt.Println("1")
}

func main() {
	deferWrapper()

	defer fmt.Println("4")

	fmt.Println("3")
}
