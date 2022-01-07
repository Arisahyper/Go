package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n >= 42 {
		fmt.Printf("AGC%03d\n", n+1)
		os.Exit(0)
	}
	fmt.Printf("AGC%03d\n", n)
}

// nを3桁になるようにゼロ埋め
