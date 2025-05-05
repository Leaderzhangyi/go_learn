package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
	Name() string
}

type Rectangle struct {
	width  float64
	height float64
	name   string // 添加 Name 字段
}

type Circle struct {
	radius float64
	name   string // 添加 Name 字段
}

type Triangle struct {
	SideLength float64
	name       string // 添加 Name 字段
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
func (r Rectangle) Name() string { // 实现 Name() 方法
	return r.name
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}
func (c Circle) Name() string { // 实现 Name() 方法
	return c.name
}

func (t Triangle) Area() float64 {
	return 0.5 * t.SideLength * t.SideLength
}
func (t Triangle) Perimeter() float64 {
	return 3 * t.SideLength
}
func (t Triangle) Name() string { // 实现 Name() 方法
	return t.name
}

func PrintShapeInfo(s Shape) {
	fmt.Println("形状:", s.Name())
	fmt.Println("面积:", s.Area())
	fmt.Println("周长:", s.Perimeter())
	fmt.Println("---")
}

func main() {
	r := Rectangle{width: 10, height: 5, name: "矩形"}
	c := Circle{radius: 3, name: "圆形"}
	t := Triangle{SideLength: 4, name: "等边三角形"}
	shapes := []Shape{r, c, t}
	for _, s := range shapes {
		PrintShapeInfo(s)
	}
}
