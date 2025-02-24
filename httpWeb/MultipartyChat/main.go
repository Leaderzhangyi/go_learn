package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var (
	clients    = make(map[net.Conn]string) // 客户端列表
	clientsMux sync.Mutex                  // 互斥锁
)

func main() {
	// 监听本地端口
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("监听端口失败:", err)
		os.Exit(1)
	}
	defer ln.Close()

	fmt.Println("聊天室已启动，等待连接...")

	// 接受客户端连接
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("连接失败:", err)
			continue
		}
		clientsMux.Lock()
		clients[conn] = conn.RemoteAddr().String() // 记录连接的客户端
		clientsMux.Unlock()

		// 启动新协程处理客户端消息
		go handleConnection(conn)
	}
}

// 处理客户端连接
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 欢迎消息
	conn.Write([]byte("欢迎加入聊天室! 请输入您的名字:\n"))

	// 获取客户端名字
	var name string
	fmt.Fscanf(conn, "%s\n", &name)
	clientsMux.Lock()
	clients[conn] = name
	clientsMux.Unlock()

	// 广播新用户加入
	broadcast(fmt.Sprintf("%s 加入了聊天室\n", name))

	// 读取客户端消息
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取消息失败:", err)
			break
		}

		// 处理GBK编码
		decodedMsg, err := decodeGBK(buf[:n])
		if err != nil {
			fmt.Println("解码失败:", err)
			continue
		}

		// 如果用户退出
		if strings.TrimSpace(decodedMsg) == "/exit" {
			break
		}

		// 广播用户消息
		broadcast(fmt.Sprintf("%s: %s", name, decodedMsg))
	}

	// 客户端断开连接时移除
	clientsMux.Lock()
	delete(clients, conn)
	clientsMux.Unlock()

	// 广播用户退出
	broadcast(fmt.Sprintf("%s 离开了聊天室\n", name))
}

// 广播消息给所有连接的客户端
func broadcast(message string) {
	clientsMux.Lock()
	defer clientsMux.Unlock()

	// 处理GBK编码
	encodedMsg, _ := encodeGBK(message)

	for conn := range clients {
		conn.Write(encodedMsg)
	}
}

// 解码GBK
func decodeGBK(data []byte) (string, error) {
	decoder := simplifiedchinese.GBK.NewDecoder()
	decodedData, _, err := transform.Bytes(decoder, data)
	if err != nil {
		return "", err
	}
	return string(decodedData), nil
}

// 编码为GBK
func encodeGBK(data string) ([]byte, error) {
	encoder := simplifiedchinese.GBK.NewEncoder()
	return transform.Bytes(encoder, []byte(data))
}
