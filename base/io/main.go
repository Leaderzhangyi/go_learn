package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("io/test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	content := make([]byte, 100) // len 与 cap相同的话可以合二为一
	n, err := file.Read(content) // 成功读出了n个字节
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(content[:n]))
	fmt.Println(content)
	fmt.Println(content[:n])

}
