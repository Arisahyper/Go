package main

import "fmt"

// 自作エラー
type UserNotFoundError struct {
	UserName string
}

// ポインタにしないと同様のエラー2個が等しいと判断されてしまう
func (e *UserNotFoundError) Error() string {
	return e.UserName + " is not found"
}

func myFunc(name string) error {
	if name == "hoge" {
		return nil // okならerrorはないのでnilを返す
	}
	return &UserNotFoundError{UserName: name} // errorならUserNotFoundを返す
}

func main() {
	name := "hogea"
	err := myFunc(name)
	if err != nil {
		fmt.Println(err)
	}
}
