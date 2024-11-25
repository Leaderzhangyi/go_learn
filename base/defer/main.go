package main

import "fmt"

func main() {
	if true {
		defer fmt.Println(1)
	} else {
		defer fmt.Println(2)
	}
	defer fmt.Println(3)
}

// Output:
// 3
// 1   // if true is true, then 1 is printed
