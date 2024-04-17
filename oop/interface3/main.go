package main

import "github.com/k0kubun/pp"

// 接口的继承

type AIn interface {
	a()
}

type BIn interface {
	b()
}

type Cin interface {
	AIn
	BIn
	c()
}

// any type
type E interface {
}

type Stu struct {
}

func (s Stu) a() {
	pp.Println("aaaaa")
}

func (s Stu) b() {
	pp.Println("bbbb")
}
func (s Stu) c() {
	pp.Println("cc")
}

func main() {
	var s Stu
	var a Cin = s
	a.a()
	a.b()
	a.c()

	var e E = s
	pp.Println(e)

}
