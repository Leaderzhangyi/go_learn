package main

import (
	"fmt"
	"net/http"
	"tutorial/httpWeb/easysignup/utils"
)

const (
	SERVER_NAME = "my_app"
	PORT        = "9999"
)

func handleSignup(w http.ResponseWriter, r *http.Request) {
	// Get the form data
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Println("Name: ", name)
	fmt.Println("Email: ", email)
	fmt.Println("Password: ", password)

}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Get the form data
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Println("Email: ", email)
	fmt.Println("Password: ", password)
}

func main() {
	server := utils.NewSdkHttpServer(SERVER_NAME)
	server.Router("/login", handleLogin)
	server.Router("/signup", handleSignup)
	server.Start(PORT)
}
