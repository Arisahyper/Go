package main

import "fmt"

func main() {
	var r, c int
	var arr[2][2] int = [2][2]int{}
	fmt.Scan(&r, &c)
	fmt.Scan(&arr[0][0], &arr[0][1])
	fmt.Scan(&arr[1][0], &arr[1][1])

	fmt.Println(arr[r-1][c-1])
}

/*
	A11 A12
	A21 A22
*/
