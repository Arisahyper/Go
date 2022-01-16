package main

import "fmt"

func goroutine(s []int, c chan int) {
	sum := 0              // 合計値の器
	for _, v := range s { // スライスの要素を取り出す
		sum += v // 合計値を計算
	}
	c <- sum // チャネルに合計値を送る
}

func main() {
	s := []int{1, 2, 3, 4, 5} // スライス / データを用意
	c := make(chan int)       // チャネルを用意

	go goroutine(s, c) // 関数にスライスとチャネルを渡す
	x := <-c           // チャネルからxへデータを受け取る

	fmt.Println(x)
}
