package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/spf13/viper"
)

var (
	clientsLock sync.Mutex
	clients     = make(map[string]net.Conn)
)

func LoadConfig() {
	// 在加载配置之前记录日志
	log.Println("加载配置文件...")

	viper.SetConfigName("config") // 配置文件名（不需要文件扩展名）
	viper.AddConfigPath(".")      // 配置文件路径
	err := viper.ReadInConfig()

	if err != nil {
		// 如果加载配置失败，记录错误日志并退出
		log.Fatalf("无法加载配置文件: %v", err)
	}

	// 配置加载成功，记录日志
	log.Println("配置文件加载成功")
}

func main() {
	// 启动日志记录
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // 设置日志格式，包含日期、时间和文件信息

	// 加载配置
	LoadConfig()

	// 打印加载的配置项
	listenPort := viper.GetString("testChat.port")
	address := viper.GetString("testChat.host")
	log.Printf("服务器正在监听端口: %s", listenPort)
	fmt.Println(listenPort)
	ln, err := net.Listen("tcp", address+":"+string(listenPort))
	if err != nil {
		log.Fatalf("无法启动TCP监听: %v", err)
	}

	// 服务器启动成功
	log.Println("服务器启动成功,地址为:", address+":"+listenPort)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("无法接受连接: %v", err)
			continue
		}
		clientsLock.Lock()
		clients[conn.RemoteAddr().String()] = conn
		log.Printf("新客户端连接: %s", conn.RemoteAddr().String())
		clientsLock.Unlock()
		go handleClient(conn)

	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("欢迎加入聊天室! 请输入您的名字:\n"))
	var name string
	fmt.Fscanf(conn, "%s\n", &name)
	broadcast(fmt.Sprintf("%s 加入了聊天室\n", name))
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("客户端断开连接: %s", conn.RemoteAddr().String())
			break
		}
		broadcast(fmt.Sprintf("%s: %s", name, string(buf[:n])))
	}
	clientsLock.Lock()
	delete(clients, conn.RemoteAddr().String())
	clientsLock.Unlock()
	broadcast(fmt.Sprintf("%s 离开了聊天室\n", name))
}

// 广播消息给所有连接的客户端
func broadcast(message string) {
	clientsLock.Lock()
	defer clientsLock.Unlock()
	for _, conn := range clients {
		conn.Write([]byte(message))
	}
}
