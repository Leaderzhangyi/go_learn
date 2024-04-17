package main

import "github.com/k0kubun/pp"

func main() {
	// r := [...]int{0: 99, 9: 88}
	// months := [...]string{1: "January" /* ... */, 12: "December"}
	// pp.Println(r)
	// pp.Println(months)
	r := [6]int{1, 2, 3, 4, 5, 6}
	m := r
	x := r[:3]
	x[0] = 99
	pp.Println(r)
	pp.Println(m)
}
