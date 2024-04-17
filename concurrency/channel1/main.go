package main

import "github.com/k0kubun/pp"

func main() {

	// 通道是一个FIFO的队列数据结构
	inChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		inChan <- i
	}

	close(inChan)
	// 若通道不关闭，直接遍历会报 deadlock
	// 因为它会自己一直读下去
	for v := range inChan {
		pp.Println(v)
	}

}
