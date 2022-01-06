package main

import "fmt"

func main() {
	var i int = 10 // int型
	var p *int     // int型のポインタ / intポインタ型
	p = &i         // 変数pに変数iのアドレスを代入
	fmt.Println(p) // 0x10 / iのアドレスが入ってる
	var j int      // int型
	j = *p         // jにアドレスであるpの中身を代入
	fmt.Println(j) // 10 / pの中身が入ってる
}

// -- メモ --
// 実数を見たい時は「*」
// アドレスを見たい時は「&」
