package main

import (
	"flag"
	"fmt"
)

func standard() {
	flag.Parse()
	args := flag.Args() // 引数を全部取得
	fmt.Println(args)

	arg1 := flag.Arg(0) // 引数を1つ目だけ取得
	fmt.Println(arg1)

	fmt.Println(flag.Arg(0), flag.Arg(10)) // 何もなかった場合空文字列が返る
}

func typeParse() {
	var ( // 型指定 / デフォルトで入る値
		i = flag.Int("int", 0, "int flag")
		s = flag.String("str", "default", "string flag")
		b = flag.Bool("bool", false, "bool flag")
	)
	flag.Parse()
	fmt.Println(*i, *s, *b)

	// go run flag.go -int 2 -str hello -bool true
}

func main() {
	typeParse()
}
