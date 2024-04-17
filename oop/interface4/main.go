package main

import "github.com/k0kubun/pp"

// 接口的定义：定义规则、定义规范、定义某种能力：

type SayHello interface {
	sayHello()
}

// 接口的实现：定义一个结构体
// 中国人：
type Chinese struct {
}

// 实现接口的方法-->具体的实现:
func (person Chinese) sayHello() {
	pp.Println("你好！")
}

func (person Chinese) good() {
	pp.Println("你是牛马！！")
}

// 美国人
type American struct {
}

func (person American) sayHello() {
	pp.Println("Hello!")

}

type integer int

func (i integer) sayHello() {
	pp.Println("say hi + ", i)
}

// 定义一个函数：专门用来各国人打招呼的函数。接受具备SayHello接口能力的变量
// 通过上下文推断，实现了多态的功能
func greet(s SayHello) {
	c, flag := s.(Chinese)
	if flag == true {
		c.good()
	} else {
		pp.Println("恭喜你逃离了地狱！！去了新的地狱美国~！")
	}

	s.sayHello()
}

func main() {
	// 创建一个中国人
	c := Chinese{}

	// 创建一个美国人
	a := American{}

	// 打招呼
	greet(a)

	// 打招呼
	greet(c)

	var i integer = 2024
	var s SayHello = i
	s.sayHello()

}
