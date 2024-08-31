package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	// 监听
	fmt.Println("FTP Server is running...")

	fmt.Println("Listen on localhost:8989")

	listen, err := net.Listen("tcp", "localhost:8989")
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go handleConnection(conn)
	}

	// 接收请求
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	currentDir := "./"
	// 向客户端发送欢迎消息
	conn.Write([]byte("220 FTP Server Ready.\n"))
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		cmd := strings.TrimSpace(string(buf[:n])) // 去除空格 + 只读1024字节中实际used的字节数
		fmt.Println("Received command2:", cmd)
		switch {
		case strings.HasPrefix(cmd, "ls"):
			handleList(conn, currentDir)
		}
	}
}

func handleList(conn net.Conn, currentDir string) {
	files, err := os.ReadDir(currentDir)
	if err != nil {
		conn.Write([]byte("550 No such file or directory.\n"))
	}
	for _, file := range files {
		fmt.Println(file.Name())
		// conn.Write([]byte(file.Name() + "\n"))
	}
	conn.Write([]byte("226 List completed.\n"))
}
