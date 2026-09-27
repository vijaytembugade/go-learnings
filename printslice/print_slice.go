package printslice

import "fmt"

// PrintSlice must be capitalized to be usable from other packages.
func PrintSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
