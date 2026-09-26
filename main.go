package main

import (
	"fmt"
)

/*
	Type of datatypes in GO
	Basic Type:  Number, string, bool
	Aggregate Type : Array, Struct
	Refrence Type : Pointers, Slices, functions , Channel, Maps
	Interfaces


*/

// basic data types
var (
	ToBe      bool   = false
	MaxInt    uint64 = 1<<64 - 1
	integer8  int8   = 21
	integer32 int32  = 238879827
	interger  int    = 134989839489238984
	str       string = "hello"
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", integer32, integer32)
	fmt.Printf("Type: %T Value: %v\n", interger, interger)
	fmt.Printf("Type: %T Value: %v\n", str, str)
}
