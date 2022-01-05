package main

import "fmt"

func main() {
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	var ans int

	for _, value := range m {
		ans += value
	}

	fmt.Println(ans)
}
