package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func returnWithSquare(value int) (int, int) {
	return value, value * value
}

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	sum := add(3, 34)
	fmt.Println(sum)

	num, square := returnWithSquare(12)
	fmt.Println(num, square)
	fmt.Println(returnWithSquare(3))
	fmt.Println(returnWithSquare(89))
	fmt.Println(returnWithSquare(11))

	second, first := swap("hello", "world")
	fmt.Println(second, first)

}
