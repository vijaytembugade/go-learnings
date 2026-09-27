package main

import (
	"fmt"
)

/*
Append in slice

-> everytime append operation execute and length and capacity is not enough, the slice will be copied to a
new memory location with double the capacity.
-> with help of make we can preassign the values of capacity if we know as third parameter, otherwise it will execute as per pattern
*/

func main() {
	s := make([]int, 0)
	printSlice("s", s)

	s = append(s, 1)
	printSlice("s", s)

	s = append(s, 2)
	printSlice("s", s)

	s = append(s, 4)
	s = append(s, 6)
	s = append(s, 8)
	printSlice("s", s)

	for i := 0; i < 300; i++ {
		s = append(s, i*i)
	}

	printSlice("s", s)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
