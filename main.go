package main

import "fmt"

/*
Struct in a GO
*/

type Vertext struct {
	X int
	Y int
}

type ContectDetails struct {
	city     string
	mobileNo int
}

type UserData struct {
	name           string
	contactDetails ContectDetails
	age            int
}

func main() {
	var v = Vertext{2, 4}
	fmt.Println(v) // 2,4

	// we can update the struct values too
	v.X = 45
	fmt.Println(v) // 45,4

	user := UserData{
		name: "vijay",
		contactDetails: ContectDetails{
			city:     "pune",
			mobileNo: 1234567890,
		},
		age: 33,
	}
	fmt.Println(user)

	user.contactDetails.mobileNo = 273874982379
	fmt.Println(user)
}
