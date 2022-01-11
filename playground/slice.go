package main

import "fmt"

func main() {
	var h []string
	h = append(h, "hoge")
	fmt.Println(h)

	s := []string{"hoge", "fuga", "piyo"}
	fmt.Println(s)
	fmt.Println(len(s)) // length
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

	// slice delete
	list3 := []string{"1", "2", "3", "4", "5"}
	list3 = append(list3[:1], list3[2:]...) // 1個目の要素と3個目以降の要素を足すことで削除をする
	fmt.Println(list3)
}

// スライス slice
// 操作 コピー
