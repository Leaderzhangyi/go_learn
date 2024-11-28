package main

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func genSessionID(ctx *gin.Context) string {
	// Generate a random sessionID
	id := base64.StdEncoding.EncodeToString([]byte(ctx.Request.RemoteAddr))
	return id
}

func login(ctx *gin.Context) {
	// 生成sessionID
}

func main() {
	// Set a cookie with name "username" and value "admin"
	engine := gin.Default()
	engine.LoadHTMLFiles("login.html")
	engine.GET("/go_to_login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})

	})
	engine.POST("/login", login)

	fmt.Println("Server is running on 127.0.0.1:10000")
	engine.Run("127.0.0.1:10000")

}
