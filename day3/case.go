package main

import "fmt"

func getOsName() string {
	return "mac"
}

func main() {
	switch getOsName() {
	case "mac":
		fmt.Println("mac")
	case "win":
		fmt.Println("win")
	default:
		fmt.Println("other")
	}
}
