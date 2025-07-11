package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func worker(ctx context.Context, taskID int, ch chan<- string) {
	delay := time.Duration(rand.Intn(4)+1) * time.Second
	select {
	case <-ctx.Done():
		fmt.Printf("worker %d done\n", taskID)
	case <-time.After(delay):
		fmt.Printf("worker %d sleep %s\n", taskID, delay)
		ch <- fmt.Sprintf("worker %d result", taskID)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	ourchan := make(chan string, 3)
	for i := 1; i <= 3; i++ {
		go worker(ctx, i, ourchan)
	}

	res := []string{}
	for i := 1; i <= 3; i++ {
		select {
		case tmp := <-ourchan:
			res = append(res, tmp)
		case <-ctx.Done():
			fmt.Println("超时啦！！！！", ctx.Err())
			for _, r := range res {
				fmt.Println(r)
			}
			return
		}

	}
	fmt.Println("\n主任务成功完成，收集结果如下：")
	for _, r := range res {
		fmt.Println(r)
	}
}
