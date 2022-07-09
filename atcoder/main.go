package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(strconv.Itoa(n)[1:])
}
