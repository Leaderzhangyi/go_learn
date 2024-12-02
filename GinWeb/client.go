package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	data := "Username=zhangyi&Password=123456&Email=zhangyi@163.com&Age=25"
	resp, err := http.Get("http://localhost:5678/register?" + data)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		return
	}

	// 输出响应内容
	fmt.Println(string(body))

}
