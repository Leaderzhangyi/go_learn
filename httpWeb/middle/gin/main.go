package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func timeMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		c.Next()
		timeEplase := time.Since(now)
		fmt.Println("Time Eplase: ", timeEplase)
	}
}

func main() {
	router := gin.Default()
	router.Use(timeMW())
	router.GET("/boy", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	fmt.Println("Server is running on port 10000")
	router.Run("127.0.0.1:10001")
}
