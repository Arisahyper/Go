package main

import "fmt"

func main() {
	var n *int = new(int) // ポインタを作成 / アドレスを先に確保する
	fmt.Println(n)        // 0x0
	fmt.Println(*n)       // 0
	*n++                  // 足してみる
	fmt.Println(*n)       // 1

	var o *int
	fmt.Println(o) 		// nil
	// *o++           // panic: runtime error: invalid memory address or nil pointer dereference
	// fmt.Println(o) // nil
}
