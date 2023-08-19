package main

import (
	"github.com/k0kubun/pp/v3"
	"time"
)

func say(id string) {
	time.Sleep(time.Second)
	pp.Println("I am done! Id is" + id)
}

func main() {
	go say("1")
	go say("2")
	pp.Println("All Done!")
	return
}
