package main

import (
	"fmt"
	"time"
)

func normal(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go normal("hoge")
	normal("fuga")
}

// goroutineが終わらなくてもプログラムが終わったら終わり
