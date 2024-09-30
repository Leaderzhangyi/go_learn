package main

import (
	"fmt"
	"time"
)

const (
	DATE = "2006-01-02"
	TIME = "2006-01-02 15:04:05"
)

func main() {
	t1 := time.Now().UTC()
	fmt.Println("The current time is", t1)
	fmt.Println(t1.Unix()) // 时间戳

	location, _ := time.LoadLocation("Asia/Shanghai")
	now := t1.In(location)
	fmt.Println("The current time in Shanghai is", now.Year(), "-", now.Month(), "-", now.Day(), " ", now.Hour(), ":", now.Minute(), ":", now.Second())
	fmt.Println(now.Format(DATE))
	fmt.Println(now.Format(TIME))

	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("The time difference is", diff)
}
