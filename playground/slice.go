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

	// slice append
	list1 := []string{"hoge", "fuga"}
	list2 := []string{"piyo", "uga"}
	sumList := append(list1, list2...)
	fmt.Println(sumList)
}


// スライス slice
// 操作 コピー
