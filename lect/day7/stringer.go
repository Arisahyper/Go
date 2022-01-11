package main

import "fmt"

type Person struct {
	name string
	age  int
}

// これがあると出力が変わる
func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.name, p.age)
}

func main() {
	mike := Person{"Mike", 20}
	fmt.Println(mike)
}
