package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义请求参数结构体
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"` // 用户名，必填，长度3-20
	Password string `json:"password" binding:"required,min=6"`        // 密码，必填，最小长度6
	Email    string `json:"email" binding:"required,email"`           // 邮箱，必填，需要符合邮箱格式
	Age      int    `json:"age" binding:"required,gte=18,lte=100"`    // 年龄，必填，范围18-100
}

func main() {
	engine := gin.Default()

	engine.GET("/register", func(ctx *gin.Context) {
		var req RegisterRequest
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, req)
		}
	})
	engine.Run("0.0.0.0:5678")

}
