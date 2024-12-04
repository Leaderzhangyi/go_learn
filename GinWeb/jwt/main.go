package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func login(ctx *gin.Context) {
	privatekey := []byte("TEST")
	claims := MyClaims{
		"admin",
		jwt.RegisteredClaims{
			Issuer:    "zy",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 一天后过期
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signtoken, err := token.SignedString(privatekey)

	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.String(http.StatusOK, signtoken)
}

func postVideo(ctx *gin.Context) {
	ctx.String(http.StatusOK, "post video")

}

func jwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqToken := ctx.GetHeader("Authorization")
		if reqToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{})
			ctx.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(reqToken, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("TEST"), nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{})
			ctx.Abort()
			return
		}
		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{})
			ctx.Abort()
			return
		}
	}
}

func main() {
	engine := gin.Default()
	engine.GET("/login", login)
	engine.GET("/postVideo", jwtAuthMiddleware(), postVideo)
	engine.Run(":10101")
	fmt.Println("Server started on address: 127.0.0.1:10101")

}
