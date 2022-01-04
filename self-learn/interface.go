package main

import "fmt"

type Human struct {
	name        string
	age         int
	phoneNumber string
}

func (human Human) getName() string {
	return human.name
} 

func main() {
	var tom Human
	tom.name, tom.age, tom.phoneNumber = "Tom", 25, "123456789"
	fmt.Println(tom.getName())
}
