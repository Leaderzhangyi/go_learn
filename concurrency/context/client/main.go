package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/k0kubun/pp"
)

func main() {

	var (
		ctx  context.Context
		err  error
		r    *http.Request
		resp *http.Response
	)
	ctx, _ = context.WithTimeout(context.Background(), 9*time.Second)
	// defer cancel()
	r, err = http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"http://localhost:10005",
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err = http.DefaultClient.Do(r)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	pp.Println(string(b))
}
