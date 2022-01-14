package main

import (
	"fmt"
	"sync"
	"time"
)

func normal(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func gor(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go gor("go routine", &wg)
	normal("normal")

	wg.Wait()
}

// goroutineが終わらなくてもプログラムが終わったら終わり
