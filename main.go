package main

import "fmt"

/*
Struct and a methods
*/

type Person struct {
	firstName string
	lastName  string
	length    int
}

// Following is a method attached to a struct
// Following is a value receiver
func (person Person) getFullName() string {
	return person.firstName + " " + person.lastName
}

// pointer reciever
// Pointer reciever updates the values in main struct
func (person *Person) AddAge() {
	person.length = len(person.firstName + person.lastName)
}

func main() {
	var person = Person{
		firstName: "vijay",
		lastName:  "tembugade",
	}

	fmt.Println(person.getFullName())
	person.AddAge()
	fmt.Println(person)
}

/**
when to pointer reciever
 -  mutation to be reflected on original instance
 - If it is very large struct


 Good practise:  If you make on receiver pointer, then make all all receiver pointer

*/
