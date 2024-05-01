package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	formData := url.Values{
		"username": {"myusername"},
		"email":    {"myemail@example.com"},
		"password": {"mypassword"},
	}

	resp, err := http.PostForm("http://localhost:9999/signup", formData)
	fmt.Println("Response status:", resp.Status)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()
	resp, err = http.PostForm("http://localhost:9999/login", formData)
	fmt.Println("Response status:", resp.Status)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

}
