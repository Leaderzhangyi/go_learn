package main

import (
	"strconv"

	"github.com/k0kubun/pp"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Say() {
	pp.Println(strconv.Itoa(p.Age) + "  " + p.Name)
}

func main() {
	var per Person
	per.Age = 12
	per.Name = "燚"
	per.Say()
}
