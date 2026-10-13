package main

import "fmt"

/*
Interfaces are named collections of method signatures.

A struct can implement an interface by implementing all the methods in the interface.

*/

// Interface Definition
type Animal interface {
	Speak() string
}

// Struct Definition
type Dog struct {
	Name string
}

// Method Implementation
func (d *Dog) Speak() string {
	return "Woof, my name is " + d.Name
}

func main() {
	var a Animal = &Dog{Name: "Rex"}
	fmt.Println(a.Speak())
	exampleEmptyInterface()
}

func exampleEmptyInterface() {
	var a interface{} = 1
	fmt.Println(a)

	a = "Hello"

	fmt.Println(a)

	a = true
	fmt.Println(a)

	a = 1.0
	fmt.Println(a)

}
