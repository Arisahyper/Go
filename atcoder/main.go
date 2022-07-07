package main

import (
	"fmt"
	"strconv"
)

func mod(k, n int) int {
	if k >= n {
		k -= n
		k = mod(k, n)
	} 
	return k
}

func main() {
	var k int
	var n string
	fmt.Scanf("%d", &k)

	if mod(k, 60) >= 10 {
		n = strconv.Itoa(mod(k, 60))
	} else {
		n = "0" + strconv.Itoa(mod(k, 60))
	}

	fmt.Printf("%s:%s", strconv.Itoa(21+(k/60)), n)
}

// https://atcoder.jp/contests/abc258/tasks/abc258_a
