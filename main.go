package main

import "fmt"

/*
Struct and pointer are being used for refrencing purpose
*/

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

	user := UserData{
		name: "vijay",
		contactDetails: ContectDetails{
			city:     "pune",
			mobileNo: 1234567890,
		},
		age: 33,
	}
	fmt.Println(user)

	p := &user
	fmt.Println(p)
	p.name = "Vijay T"
	p.contactDetails.mobileNo = 98239

	fmt.Println(user)

	if &p.contactDetails == &user.contactDetails {
		fmt.Println(&p.contactDetails)
	}

	i := 1
	k := &i
	fmt.Println(*k)

}

/*
pointer to a struct → p.field (auto-deref)
pointer to a plain value (int, string, etc.) → *k
*/
