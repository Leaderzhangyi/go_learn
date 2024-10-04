package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 100)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 10; i < 20; i++ {
			ch <- i
		}
	}()

	zs := make(chan struct{}, 0)
	// 消费者
	go func() {
		for i := 0; i < 20; i++ {
			v, ok := <-ch
			if !ok {
				break
			} else {
				fmt.Println(v)
			}
		}
		zs <- struct{}{} // 通知消费者结束
	}()

	wg.Wait()
	close(ch)
	<-zs // 从zs管道读数据，等待消费者结束
}
