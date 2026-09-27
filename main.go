package main

import (
	"fmt"
	"go-learning/printslice"
)

var arr = []int{1, 2, 3, 4, 5, 6, 7, 9}

func main() {
	sqArr := make([]int, 0, len(arr))

	for index, value := range arr {
		fmt.Println(index)
		sqArr = append(sqArr, value*value)
	}

	printslice.PrintSlice("sqArr", sqArr)
}
