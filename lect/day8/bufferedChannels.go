package main

import "fmt"

func main() {
	ch := make(chan int, 2) // チャネルを用意
	ch <- 100               // チャネルにデータを送る
	fmt.Println(len(ch))    // チャネルの要素数を表示
}
