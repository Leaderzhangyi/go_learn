package main

import (
	"log"
	"net/http"
	"time"

	"github.com/k0kubun/pp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-time.After(1 * time.Second):
		pp.Println("Finish...")
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			pp.Println(err.Error())
		}
	}
}

func main() {

	http.HandleFunc("/", handler)
	// http.Handle()
	log.Fatalln(http.ListenAndServe(":10005", nil))
}
