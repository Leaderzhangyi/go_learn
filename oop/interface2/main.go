package main

import "github.com/k0kubun/pp"

// 接口的定义：定义规则、定义规范、定义某种能力：

type AIn interface {
	a()
}

type BIn interface {
	b()
}

type Stu struct {
}

func (s Stu) a() {
	pp.Println("aaa")
}

func (s Stu) b() {
	pp.Println("bbb")

}

func main() {
	var (
		s Stu
		a AIn = s
		b BIn = s
	)

	a.a()
	b.b()
}
