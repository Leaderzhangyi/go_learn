package main

import (
	"fmt"
	"net/http"
	"time"
)

type middle func(http.Handler) http.Handler

type Router struct {
	mux         map[string]http.Handler
	middleChain []middle
}

func NewRouter() *Router {
	r := &Router{
		mux:         make(map[string]http.Handler, 10),
		middleChain: make([]middle, 0, 10),
	}
	return r
}

func (r *Router) Use(m middle) {
	r.middleChain = append(r.middleChain, m)
}

func (r *Router) Add(path string, handler http.Handler) {
	var mergedHandler = handler
	for i := len(r.middleChain) - 1; i >= 0; i-- {
		mergedHandler = r.middleChain[i](mergedHandler)
	}
	r.mux[path] = mergedHandler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, ok := r.mux[req.URL.Path]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler.ServeHTTP(w, req)
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		timeEplase := time.Since(now)
		fmt.Println("Time Eplase:", timeEplase.Microseconds())
	})
}

func boy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("boy"))
}

func main() {
	router := NewRouter()
	router.Use(timeMiddleware)
	router.Add("/test", http.HandlerFunc(boy))
	fmt.Println("Server is running on :10000")
	http.ListenAndServe(":10000", router)
}
