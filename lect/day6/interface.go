package main

import "fmt"

type Human interface {
	Say()
	Rename()
}

type Person struct {
	name string
}

func (p Person) Say() {
	fmt.Println(p.name)
}

func (p *Person) Rename(){
	p.name = "Mr." + p.name
	fmt.Println(p.name)
}

func main() {
	var mike Human = &Person{"Mike"}
	mike.Say()
	mike.Rename()
}
