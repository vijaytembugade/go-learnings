package main

import "fmt"

/*
	map is a collection of key-value pairs
	key is unique and value is not unique
	value can be of any type

	How is is different from struct?
	1. map is a collection of key-value pairs
	2. struct is a collection of fields / statically typed
*/

type Address struct {
	city    string
	pincode int
}

type User struct {
	name    string
	age     int
	address Address
}

func main() {
	var user = make(map[int]User)
	fmt.Println(user)

	for i := 0; i < 10; i++ {
		user[i] = User{
			name: "Vijay",
			age:  28 * i,
			address: Address{
				city:    "pune",
				pincode: 411014,
			},
		}
	}

	var address = map[string]Address{
		"pune": {city: "pune"},
	}

	city, ok := address["pune"]
	if ok && city.pincode != 0 {
		fmt.Println("pincode present:", city.pincode)
	} else {
		fmt.Println("pincode not present")
	}

	fmt.Println(user)
	fmt.Println(address)
	fmt.Println(city, ok)

	for key, value := range user {
		fmt.Println(key, value)
	}

}
