package main

import "fmt"

func connectDB() {
	panic("DB PANIC")
}

func saveDB() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	connectDB()
}

func main() {
	saveDB()
	fmt.Println("Hello World")
}

// panicは言語レベルで非推奨
// エラーハンドリングを使おう
