package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	format := "2006-01-02 15:04:05"
	fmt.Println(now.Format(format))
}
