package main

import "fmt"

func main() {
	var n int = 100
	fmt.Println(n)  // 100
	fmt.Println(&n) // 変数nのメモリのアドレス

	var p *int = &n // ポインタ型
	fmt.Println(p)  // 変数nのメモリのアドレス / 上と同じ
	fmt.Println(*p) // 変数nの値 / 変数nの値を参照する
}

func p(arg *int) int { // メモリの番地を渡す
	return *arg + 1
}

// 番地指定したい時はポインタ型で渡す
// 参照したい時は、&をつける
