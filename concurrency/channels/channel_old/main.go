package main

import (
	"github.com/k0kubun/pp/v3"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(id string) {
	time.Sleep(time.Second)
	pp.Println("I am done! Id is" + id)
	wg.Done() //task is Done!
}

func main1() {
	wg.Add(2) // 共两个任务
	//go say("1")
	go func(id string) {
		pp.Println("First!!!" + id)
		wg.Done()
	}("1")
	go say("2")
	wg.Wait() //等待所有任务结束，done相当于--(类pv操作)
	pp.Println("All Done!")
	return
}
