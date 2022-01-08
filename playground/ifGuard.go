package main

import "fmt"

func numberDiscrimination(i int) string {
	if i == 100 {
		return "i is 100"
	}
	if i == 200 {
		return "i is 200"
	}
	if i == 300 {
		return "i is 300"
	}
	return "i is not 100, 200 or 300"
}

func main() {
	var i int = 101
	result := numberDiscrimination(i)
	fmt.Println(result)
}

// ガード節？
