package main

import (
	"sync"
	"time"

	"github.com/k0kubun/pp"
)

var totNum int
var lock sync.RWMutex
var wg sync.WaitGroup

func read() {

	defer wg.Done()
	lock.RLock()
	pp.Println("starting read csv...")
	time.Sleep(500 * time.Millisecond)
	pp.Println("read successful")
	lock.RUnlock()

}

func write() {

	defer wg.Done()
	lock.Lock()
	pp.Println("starting write csv...")
	time.Sleep(2 * time.Second)
	pp.Println("write successful")
	lock.Unlock()

}

func main() {
	wg.Add(6)
	for i := 0; i < 5; i++ {
		go read()
	}
	go write()
	wg.Wait()
}
