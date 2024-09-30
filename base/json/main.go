package main

import (
	"encoding/json"
	"fmt"
)

// 磁盘与网络信息传输，本质是一个二进制流

type Student struct {
	Name   string
	Age    int
	Gender string
}

type Class struct {
	Id       string
	Students []Student
}

func main() {
	s := Student{"Tom", 18, "Male"}
	class := Class{
		Id:       "1001",
		Students: []Student{s, s, s},
	}

	// 序列化为json字符串
	data, err := json.Marshal(class)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	var c2 Class
	err = json.Unmarshal(data, &c2) // 反序列化为结构体
	//tips: go的函数传值采用的是值拷贝，所以这里要传指针，不然不会改变原来的值
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c2)
	// sonic 字节跳动的json库 比官方的json库更快
}
