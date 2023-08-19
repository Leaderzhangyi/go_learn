package main

import (
	"github.com/k0kubun/pp/v3"
	"sync"
	"sync/atomic"
)

var counter int32
var wg sync.WaitGroup
var mtx sync.Mutex
var ch = make(chan int32, 1)

func UnsafeCounter() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		counter += 1
	}
}

func MtxSafeCounter() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		mtx.Lock()
		counter += 1 // 竞争资源
		mtx.Unlock()
	}
}

func AtomicSafeCounter() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&counter, 1)
	}
}

func ChSafeCounter() {
	defer wg.Done()
	for i := 0; i < 20000000; i++ {
		count := <-ch
		count++
		ch <- count
	}
}

func main() {
	wg.Add(2)
	go ChSafeCounter()
	go ChSafeCounter()
	ch <- 0
	wg.Wait()
	pp.Println("counter = ", <-ch)
}
