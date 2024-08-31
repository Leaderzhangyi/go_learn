package main

import (
	"fmt"
	"time"
)

// import (
// 	"sync"

// 	"github.com/k0kubun/pp"
// )

// func other(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 10; i++ {
// 		pp.Println("I'm Lin Yingjie")
// 		// time.Sleep(1 * time.Second)
// 	}
// }

// func main() {
// 	wg := sync.WaitGroup{}
// 	wg.Add(1)
// 	go other(&wg)
// 	for i := 0; i < 10; i++ {
// 		pp.Println("I'm Zhang Yi")
// 		// time.Sleep(1 * time.Second)
// 	}
// 	wg.Wait()
// }

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
