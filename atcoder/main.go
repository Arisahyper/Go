package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	for {
		if len(s) >= 6 {
			break
		}
		s += s
	}

	fmt.Println(s[0:6])
}
