package main

import "fmt"

/*
Arrays in go

In Go , size of array is part of array
eg. var a [2]string
*/

type User struct {
	name string
	age  int
}

func main() {
	var a [3]string
	a[0] = "hello"
	a[1] = "world"
	a[2] = "damn"

	fmt.Println(a)

	str := [3]string{"jivan", "saral", "hai"}
	fmt.Println(str)

	var primes = [4]int{2, 3, 5, 7}
	fmt.Println(primes)

	var users = [2]User{{
		name: "some",
		age:  23,
	}, {
		age:  24,
		name: "jwepr",
	}}

	fmt.Println(users)
}
