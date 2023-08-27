package main

import (
	"Go_pro/downloader"
	"github.com/k0kubun/pp/v3"
)

func main() {
	request := downloader.InfoRequest{Bvids: []string{"BV1Ff4y187q9", "BV1vf4y1Q7An"}}
	info, err := downloader.BatchDownloadVideoInfo(request)
	if err != nil {
		panic(err)
	}
	pp.Println(info)
}
