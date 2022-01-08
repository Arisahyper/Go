package main

import "fmt"

type Human interface {
	Say() 	 string
}

type Person struct {
	name string
}

func (p Person) Say() string {
	return p.name
}

func DriveCar(human Human) {
	if human.Say() == "Mike" {
		fmt.Println("Drive a car")
	} else {
		fmt.Println("Don't drive a car")
	}
}

func main() {
	var mike Human = Person{"Mike"}
	var john Human = Person{"John"}

	DriveCar(mike)
	DriveCar(john)
}
