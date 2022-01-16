package main

import "fmt"

func goroutine(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5} // スライス / データを用意
	c := make(chan int)       // チャネルを用意

	go goroutine(s, c) // 関数にスライスとチャネルを渡す
	x := <-c           // チャネルからxへデータを受け取る

	fmt.Println(x)
}
