package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()	
	args := flag.Args()	// 引数を全部取得
	fmt.Println(args)

	arg1 := flag.Arg(0)	// 引数を1つ目だけ取得
	fmt.Println(arg1)

	fmt.Println(flag.Arg(0), flag.Arg(10))	// 何もなかった場合空文字列が返る
}
