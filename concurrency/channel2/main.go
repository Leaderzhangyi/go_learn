package main

import (
	"sync"

	"github.com/k0kubun/pp"
)

var wg sync.WaitGroup

func writeData(opChan chan int) {
	defer wg.Done()
	for i := 1; i <= 50; i++ {
		opChan <- i
		pp.Println("存：", i)

	}
	close(opChan)
}

func readData(opChan chan int) {
	defer wg.Done()
	for v := range opChan {
		pp.Println("取：", v)
	}
}

func main() {
	intChan := make(chan int, 50)
	wg.Add(2)
	go writeData(intChan)
	go readData(intChan)
	wg.Wait()
	pp.Println("完成....")
}
