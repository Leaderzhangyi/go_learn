package main

import (
	"fmt"
	"net/http"
	"time"
)

// 在全局作用域中，Go要求使用 var 关键字来声明变量，以确保变量的定义是显式的、可读的，并且类型清晰。这有助于维护代码的可读性和防止意外行为。
var limitchannel = make(chan struct{}, 10)

func boy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("boy"))
}

func total(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world! \n This is a total function."))
}

func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limitchannel <- struct{}{}
		next.ServeHTTP(w, r)
		<-limitchannel
	})
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		timeEplase := time.Since(now)
		fmt.Println("Time Eplase:", timeEplase.Microseconds())
	})
}

func main() {
	http.Handle("/boy", timeMiddleware(http.HandlerFunc(boy)))
	http.HandleFunc("/", total)
	fmt.Println("Server is running on http://127.0.0.1:10000")
	http.ListenAndServe("127.0.0.1:10000", nil)
}
