package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func main() {
	ch := make(chan int, 2)
	s := []int{1, 2, 3, 4, 5}
	go goroutine1(s, ch)
	for c := range ch {
		fmt.Println(c)
	}
}
