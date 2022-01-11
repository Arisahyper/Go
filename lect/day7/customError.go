package main

import "fmt"

// 自作エラー
type UserNotFound struct {
}

func myFunc() error {
	return nil
}

func main() {
	err := myFunc()
	if err != nil {
		fmt.Println(err)
	}
}
