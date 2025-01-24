package main

import (
	"fmt"
)

func testDefer() int {
	i := 1
	defer func() {
		i++
	}()
	defer func() {
		i++
	}()
	return i
}

func main() {
	if true {
		defer fmt.Println(1)
	} else {
		defer fmt.Println(2)
	}
	defer fmt.Println(3)
	defer fmt.Println(4)
	defer fmt.Println(5)
	// fmt.Println(testDefer())
}

// Output:
// 3
// 1   // if true is true, then 1 is printed
