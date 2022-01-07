package main

import (
	"fmt"
	"os"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var ans int

	if x > y {
		ans = 0
		fmt.Println(ans)
		os.Exit(0)
	}

	if (y-x)%10 != 0 {
		ans = (y-x)/10 + 1
		fmt.Println(ans)
		os.Exit(0)
	}
	ans = (y-x)/10
	
	fmt.Println(ans)
}
