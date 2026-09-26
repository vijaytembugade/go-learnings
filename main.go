package main

import (
	"fmt"
)

var var1, var2, sum int // in go variable is always initized with 0 state.
var isCorrect bool      // boolen is initilised with false value

func main() {
	i := 0 // this type of declaration only valid inside a function
	fmt.Println(isCorrect)
	fmt.Println(var2, i)

	var j = 18
	fmt.Println(j)

	var k string = "people"
	fmt.Println(k)
}
