package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wk sync.WaitGroup

func player(name string, ch chan int) {
	defer wk.Done() //task is Done!
	for {
		ball, ok := <-ch // 返回两个量，第一个为通道的值，第二个为bool变量
		if !ok {         // 通道关闭
			fmt.Printf("channel is closed! %s wins!\n", name)
			return
		}

		n := rand.Intn(100)
		if n%10 == 0 {
			// 十分之一机率没接上，就关闭通道

			close(ch)
			fmt.Printf("%s missed the ball.... %s loss!\n", name, name)
			return
		}
		ball++
		fmt.Printf("%s receives ball %d\n", name, ball)
		ch <- ball
	}
}

func main() {
	wk.Add(2)
	// 利用通道模拟打球
	ch := make(chan int, 0) //no缓存通道
	go player("zy", ch)
	go player("ljw", ch)
	ch <- 0   // 必须定义了goroutine才能传球，不然会一直卡在这里
	wk.Wait() //等待所有任务结束，done相当于--(类pv操作)
	return
}
