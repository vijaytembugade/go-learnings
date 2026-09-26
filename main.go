package main

import "fmt"

// type conversion in a go is not implecite but it is explicite.
// It should be done properly otherwise it will give runtime erros
func main() {
	var i int8 = 3
	var j int32 = int32(i)
	fmt.Println(j)
}
