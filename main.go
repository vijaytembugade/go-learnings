package main

import "fmt"

/*
reference types in a go ,
there are 5 types of reference types:
pointers, slices, maps, functions, channels

& -> gives you the address of the variable (address of operator)
* -> gives you the value of the variable (dereferencing)
*/
func main() {
	pointerExample()
}

func pointerExample() {
	i := 21

	p := &i         // p is address of i
	fmt.Println(p)  // print the address of i -> 0x52e0d216020
	fmt.Println(*p) // print the value of i -> 21

	var k int = 34
	fmt.Println(k) // 34
	pointerToK := &k
	*pointerToK = 56 // it changes the value of k, because pointerToK is poiting to its value not address
	fmt.Println(k)   // 56
}
