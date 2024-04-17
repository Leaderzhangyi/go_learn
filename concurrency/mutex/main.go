package main

import (
	"fmt"
	"sync"

	"github.com/k0kubun/pp"
)

var totalNum int
var wg sync.WaitGroup
var lock sync.Mutex

func Add() {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		lock.Lock()
		totalNum += 1
		lock.Unlock()
	}
}

func Reduce() {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		lock.Lock()
		totalNum -= 1
		lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go Add()
	go Reduce()
	wg.Wait()
	pp.Println(totalNum)
	fmt.Printf("9 = %d", totalNum)
}
