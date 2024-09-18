package main

import (
	"fmt"
)

func main() {
	s := "Hello, world!"
	// 字符串的本质是不可修改的byte数组
	fmt.Println(s)
	ss := "你好，世界！" // 一个中文占3个byte
	fmt.Printf("org string len = %d\n", len(ss))
	arr := []rune(ss)
	fmt.Printf("rune array len = %d\n", len(arr))

	for i, r := range ss {
		fmt.Printf("ss[%d] = %c ascii = %v\n", i, r, r)
	}
	fmt.Println("----------------------------------------")
	for i, r := range arr {
		fmt.Printf("arr[%d] = %c\n", i, r)
	}

	test := `dadada
	dadadada
	dadadad`
	fmt.Println(test)

}
