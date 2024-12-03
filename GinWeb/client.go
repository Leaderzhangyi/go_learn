package main

import (
	"fmt"
	"io"
	"net/http"
)

func testVaildator() {
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

func testCookies() {
	// 发送请求
	resp, err := http.Get("http://localhost:10101/login")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		return
	}
	// 输出响应内容
	fmt.Println(string(body))

	fmt.Println("获取cookie:")
	for _, cookie := range resp.Cookies() {
		fmt.Println(cookie.Name, cookie.Value)
	}

	// 发送带有cookie的请求
	req, err := http.NewRequest("GET", "http://localhost:10101/postVideo", nil)
	if err != nil {
		panic(err)
	}
	// 设置cookie
	for _, cookie := range resp.Cookies() {
		req.AddCookie(cookie)
	}
	//	发送请求
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// 读取响应内容
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		return
	}
	// 输出响应内容
	fmt.Println(string(body))

}

func main() {
	// testVaildator()
	testCookies()
}
