package main

import (
	"sync"

	"github.com/k0kubun/pp"
)

func other(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		pp.Println("I'm Lin Yingjie")
		// time.Sleep(1 * time.Second)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go other(&wg)
	for i := 0; i < 10; i++ {
		pp.Println("I'm Zhang Yi")
		// time.Sleep(1 * time.Second)
	}
	wg.Wait()
}
