package main

import "fmt"

func main() {
	os := "mac"
	switch os {
	case "mac":
		fmt.Println("mac")
	case "win":
		fmt.Println("win")
	default:
		fmt.Println("other")
	}
}
