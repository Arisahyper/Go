package main

import "fmt"

func main() {
	var i int = 100
	var j int = 200
	var p1 *int = &i
	var p2 *int = &j

	i = *p1 + *p2
	p2 = p1

	j = *p2 + i
	fmt.Println(j)
}
