package main

import "fmt"

/*
Slices in go

- Slice is always has background array which is being referenced through
- Slice has 3 parts i.e length, capacity and pointer
- Slice is refrence type
- In slice we do not have size of a items, but in array we have to have a size
- zero value of slice is nil
*/

type Person struct {
	Name string
	Age  int
}

func main() {
	var a = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceFromAnnArray(a)

	// declaring an slice without array
	q := []int{23, 35, 67, 89}
	fmt.Println(q)

	// slice of struct
	people := []Person{
		{Name: "John", Age: 20},
		{Name: "Jane", Age: 21},
		{Name: "Jim", Age: 22},
	}
	fmt.Println(people)

	// slice of slice
	sliceOfSlice := people[0:2]
	fmt.Println(sliceOfSlice)

	emptySlice()

}

func sliceFromAnnArray(arr [10]int) {
	a := arr[0:3]  // length 3, capacity : 10, poiter : 0 -> 1,2,3
	b := arr[4:10] // length 6, capacity: 6, pointer: 4  -> 4,5,6,7,8

	fmt.Println(a, b)

	b[4] = 15
	a[0] = 20
	fmt.Println(a, b) // [20 1 2] [4 5 6 7 15 9]
	fmt.Println(arr)  // [20 1 2 3 4 5 6 7 15 9]
}

func emptySlice() {
	var s = []int{2, 3, 4, 5}
	fmt.Println(s, len(s), cap(s))

	s = s[1:3]
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("slice is nil")
	}

}
