package main

import "fmt"

func standard() {
	n := []int{1, 2, 3, 4, 5}
	fmt.Println(n[1:3]) // [2 3]

	for i, v := range n[1:3] {
		fmt.Println(i, v)
	}
}

func board() {
	borad := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(borad[1][1])
	// [[1 2 3] [4 5 6]]
}

func capMake() {
	n := []int{1, 2, 3, 4, 5}
	fmt.Printf("len: %d\ncap: %d\n", len(n), cap(n))
}

func main() {
	// standard()
	// board()
	capMake()
}
