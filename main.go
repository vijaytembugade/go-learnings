package main

import "fmt"

const Pi = 3.14

func main() {
	v := 42.3

	k := v
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("v is of type %T\n", k)

	fmt.Printf("%v %T\n", Pi, Pi)
}
