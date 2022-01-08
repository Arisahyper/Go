package main

import "fmt"

type Human interface {
	Say() string
	Rename() string
}

type Person struct {
	name string
}

func (p Person) Say() string {
	fmt.Println(p.name)
	return p.name
}

func (p *Person) Rename() string {
	p.name = "Mr." + p.name
	return p.name
}

func DriveCar(h Human) {
	if h.Say() == "Mike" {
		fmt.Println("Drive a car")
	} else {
		fmt.Println("Don't drive a car")
	}
}

func main() {
	var mike Human = &Person{"Mike"}
	var john Human = &Person{"John"}

	// mike.Say()
	// mike.Rename()

	DriveCar(mike)
	DriveCar(john)
}
