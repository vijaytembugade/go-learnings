package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(12))
	fmt.Println("My favorite number is", rand.Intn(90))
	fmt.Println("My favorite number is", math.Pi)
	fmt.Println("My favorite number is", math.Sqrt(2))
}
