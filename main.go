package main

import (
	"fmt"
	"time"
	"tutorial/runner"
)

func createTask() func(int) {
	return func(id int) {
		time.Sleep(time.Second)
		fmt.Printf("Task complete #%d\n", id)
	}

}

func main() {
	//request := downloader.InfoRequest{Bvids: []string{"BV1Ff4y187q9", "BV1vf4y1Q7An"}}
	//info, err := downloader.BatchDownloadVideoInfo(request)
	//if err != nil {
	//	panic(err)
	//}
	//pp.Println(info)
	r := runner.New(4 * time.Second)
	r.AddTasks(createTask(), createTask(), createTask())
	err := r.Start()
	switch err {
	case runner.ErrInterrupt:
		fmt.Println("tasks interrupted")
	case runner.ErrTimeout:
		fmt.Println("tasks timeout")
	default:
		fmt.Println("all tasks finished!")
	}

}
