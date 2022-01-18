package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup // WaitGroupを作成

	for i := 1; i <= 5; i++ { // 5回繰り返す
		wg.Add(1) // WaitGroupに1つ追加
		i := i
		go func() {
			defer wg.Done() // Doneを呼ぶと自動的に1つ減らす
			worker(i)       // workerを実行
		}()
	}

	wg.Wait() // WaitGroupが0になるまで待つ
}

// 一個渡したら(Add)一個減らす(Done)
