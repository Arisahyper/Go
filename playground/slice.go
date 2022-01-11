package main

import "fmt"

func main() {
	var h []string
	h = append(h, "hoge")
	fmt.Println(h)

	s := []string{"hoge", "fuga", "piyo"}
	fmt.Println(s)
	fmt.Println(len(s))	// length
	fmt.Println(cap(s)) // capacity

	// slice copy
	list := []string{"hoge", "fuga", "piyo"}
	dist := make([]string, len(list))
	copy(dist, list)
	fmt.Println(dist)
}


// スライス slice
// 操作 コピー
