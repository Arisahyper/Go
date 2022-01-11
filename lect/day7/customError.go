package main

import "fmt"

// 自作エラー
type UserNotFound struct {
	UserName string
}

func (e *UserNotFound) Error() string {
	return e.UserName + " is not found"
}

func myFunc(name string) error {
	if name == "hoge" {
		return nil // okならerrorはないのでnilを返す
	}
	return &UserNotFound{name} // errorならUserNotFoundを返す
}

func main() {
	name := "hoge"
	err := myFunc(name)
	if err != nil {
		fmt.Println(err)
	}
}
