package main

import "fmt"

/*
generics are a way to write code that can be used with different types.

For example, we can write a function that can be used with different types of values.

We can use the type parameter T to represent the type of the value.

We can use the type parameter T to represent the type of the value.

We can use the type parameter T to represent the type of the value.
*/

func main() {
	exampleGenerics(1)
	exampleGenerics("Hello")
	exampleGenerics(true)
}

func exampleGenerics[T comparable](value T) {
	fmt.Println(value)
}
