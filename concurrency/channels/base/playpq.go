package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func player(name string, ch chan int) {
	defer wg.Done()
	for {
		ball, ok := <-ch
		if !ok {
			log.Println(name, "channel closed")
			return
		}
		n := rand.Intn(1000)
		if n%100 == 0 {
			close(ch)
			log.Printf("%s looses the game\n", name)
			return
		}
		mutex.Lock()
		ball++
		mutex.Unlock()
		fmt.Printf("%s throw the ball %d\n", name, ball)
		ch <- ball // 球传给对手
	}
}

func main() {

	wg.Add(2)
	ch := make(chan int)
	go player("linyingjie", ch)
	go player("zhangyi", ch)

	ch <- 0
	wg.Wait()

}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var mu sync.Mutex
// var chain string

// func main() {
// 	chain = "main"
// 	A()
// 	fmt.Println(chain)
// }
// func A() {
// 	mu.Lock()

// 	chain = chain + " --> A"
// 	mu.Unlock()
// 	B()
// }
// func B() {
// 	chain = chain + " --> B"
// 	C()
// }
// func C() {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	chain = chain + " --> C"
// }
