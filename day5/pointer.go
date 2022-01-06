package main

import "fmt"

func main() {
	var n int = 100
	fmt.Println(n)  // 100
	fmt.Println(&n) // 変数nのメモリのアドレス

	var p *int = &n
	fmt.Println(p)  // 変数nのメモリのアドレス / 上と同じ
	fmt.Println(*p) // 変数nの値 / 変数nの値を参照する
}
