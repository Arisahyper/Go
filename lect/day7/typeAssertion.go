package main

import "fmt"

// 引数 interface{} は、型を指定せず引数を取れる
func typeBranch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("bool", v)
	default:
		fmt.Println("unknown", v)
	}
}

func do(i interface{}){
	// ii := i * 2 // error
	ii := i.(int) * 2 // type assertion
	fmt.Println(ii)
}

func main() {
	do(1)
	typeBranch(1)
	typeBranch("hoge")
	typeBranch(true)
}

// type assertion
// interface{}に対して型を指定すること

// type conversion
// cast
// 一般的な標準の型変換
