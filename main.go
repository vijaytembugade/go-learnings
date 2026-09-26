package main

import "fmt"

/*
defer in go
defer push the evalution into a stack.
after function returned, go takes out stack defer expesions and evaluate it

WhY do we need defer ???
1.  if code throws an error or it panics, in that cases we need to ecute few things
which will execute a code logic, or things, we can use defer there so that we get to finish few started things
eg. If weopen a DB connection and meanwhile something happen and we did not close the connection it will cause memory leak
So, defer help to make those close connection work at the end of a function
*/

func main() {
	example()
	returenedValue := counting()
	fmt.Println(returenedValue)
}

/*
out put of above
world
hello
starting
done
9
8
7
6
5
4
3
2
1
0
true
*/

func example() {
	defer fmt.Println("hello")
	fmt.Println("world") // world will be printed first and the hello
}

func counting() bool {
	fmt.Println("starting")
	for i := range 10 {
		defer fmt.Println(i)
	}
	// [] -> in this stack, value will be pushed like i=0 -> 0, i=1 -> 1 i.e [9,8,7,6,5,4,3,2,1]
	// and poped out values from top of a stack.
	// so output of print will be [9,8,7,6,5,4,3,2,1] like this
	fmt.Println("done")
	return true
}
