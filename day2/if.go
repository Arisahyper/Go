package main

import "fmt"

func standarIf() {
	num := 10
	if num%2 == 0 {
		fmt.Println("even")
	} else if num == 0 {
		fmt.Println("num == 0")
	} else {
		fmt.Println("odd")
	}
}

func symbolIf() {
	x, y := 10, 10
	if x == 10 && y == 10 {
		print("&&")
	}
	if x == 10 || y == 10 {
		print("||")
	}
}

func main() {
	if num := 10; num%2 == 0 {
		fmt.Println("even")
	}
}
