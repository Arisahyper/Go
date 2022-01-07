package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	sl := strings.Split(s, "")

	var n1, _ = strconv.Atoi(sl[0])
	var n2, _ = strconv.Atoi(sl[2])

	fmt.Println(n1 * n2)
}
