package utils

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server interface {
	Router(patten string, handler http.HandlerFunc) error
	Start(address string) error
	Shutdown(ctx context.Context) error
}

type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Router(patten string, handler http.HandlerFunc) error {
	http.HandleFunc(patten, handler)
	fmt.Println("Router added: ", patten)

	return nil
}

func (s *sdkHttpServer) Start(address string) error {
	fmt.Println("Server is running on port:", address)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println("Failed to start server: ", err)
		return err
	}
	return nil
}

func (s *sdkHttpServer) Shutdown(ctx context.Context) error {
	time.Sleep(1 * time.Second)
	fmt.Printf("#{s.Name} shutdown...\n")
	return nil
}

func NewSdkHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}
