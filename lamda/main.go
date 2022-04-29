package main

import "fmt"

func main() {
	f := func(name string) string { return "Hello, " + name }

	fmt.Println(f("Tanaka"))
}
