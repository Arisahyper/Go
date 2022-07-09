package main

import "fmt"

func main() {
	var n, x int = 1, 3
	fmt.Scan(&n, &x)

	ch := ((x-1)/n + 'A')
	fmt.Println(string(ch))
}
