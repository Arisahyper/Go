package main

import "fmt"

func main() {
	ch := make(chan int, 2) // チャネルを用意
	ch <- 100               // チャネルにデータを送る
	ch <- 200               // チャネルにデータを送る

	close(ch) // チャネルを閉じる / rangeで取り出す際はこれが必要

	for c := range ch {
		fmt.Println(c)
	}
}
