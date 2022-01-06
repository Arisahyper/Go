package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}

	// キーのみ取得
	for k := range m {
		fmt.Println(k)
	}

	// 値のみ取得
	for _, v := range m {
		fmt.Println(v)
	}
}
