package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:17312")
	log.Println("localhost:17312 listening...  ")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handleConn(conn)

	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal(err)
	}

	for {
		_, err := io.WriteString(c, time.Now().In(location).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
