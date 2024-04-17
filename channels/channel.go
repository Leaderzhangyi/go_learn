package main

import "fmt"

func main() {
	// 定义channel
	c := make(chan int)

	go func() {
		defer fmt.Println("Go End..")
		fmt.Println("goruntime running！")
		c <- 89138109
	}()
	num := <-c
	fmt.Println("num = ", num)

	return

}
