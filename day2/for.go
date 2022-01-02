package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
		if i == 5{
			fmt.Println("continueでskip")
			continue
		}
	}
}

