package main

import (
	"fmt"
)

// go only have for loop

func main() {
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// while loop execution in go
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	fooBar(12)
	fibo(10)
	infiniteLoopUseCase(1)
}

func fooBar(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("foobar", i)
			continue
		}
		if i%3 == 0 {
			fmt.Println("foo", i)
			continue
		}
		if i%5 == 0 {
			fmt.Println("bar", i)
			continue
		}
	}
}

func fibo(n int) {
	var first = 0
	var second = 1
	fmt.Println(first, second)
	for i := 0; i < n; i++ {
		var temp = first + second
		fmt.Println(temp)
		first = second
		second = temp
	}
}

func infiniteLoopUseCase(n int) {
	for {
		if n == 100 {
			break
		}
		n = n + 1
	}
	fmt.Println(n)
}
