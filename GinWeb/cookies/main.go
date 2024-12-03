package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"sync"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Age  int    `json:"age"`
	Vip  bool   `json:"vip"`
}

var (
	userinfos sync.Map
)

const (
	authCookieName = "auth"
)

func getSessionId(ctx *gin.Context) string {
	// 生成新的SessionId
	return base64.StdEncoding.EncodeToString([]byte(ctx.Request.RemoteAddr))
}

func login(ctx *gin.Context) {
	// 从request	中获取参数 与数据库作对比 此处省略

	for _, item := range ctx.Params {
		fmt.Println(item)
	}

	session_id := getSessionId(ctx)
	user := User{
		Name: "zhangyi",
		Role: "admin",
		Age:  25,
		Vip:  true,
	}
	userinfo, _ := sonic.Marshal(user)
	userinfos.Store(session_id, userinfo)
	ctx.SetCookie(
		authCookieName, //cookies Name
		session_id,     // sessionId
		3000,           // 有效时间 单位秒
		"/",            // 作用域
		"localhost",    // 域名
		false,          // 安全
		true,           // httpOnly
	)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "login success",
	})
}

func postVideo(ctx *gin.Context) {
	ctx.String(http.StatusOK, fmt.Sprintf("%s 发布了视频", ctx.GetString("name")))
}

func comment(ctx *gin.Context) {
	ctx.String(http.StatusOK, "评论成功")
}

func checkAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// session_id := ctx.Request.Cookies()
		session_id := getSessionId(ctx)
		sessionIdExists := false
		for _, cookie := range ctx.Request.Cookies() {
			fmt.Println(cookie.Name, cookie.Value)
			if cookie.Name == authCookieName && cookie.Value == session_id {
				sessionIdExists = true
				break
			}
		}
		if !sessionIdExists {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			ctx.Abort()
		}
		if v, ok := userinfos.Load(session_id); !ok {
			ctx.String(http.StatusForbidden, "身份认证失败")
			ctx.Abort()
		} else {
			var user User
			sonic.Unmarshal(v.([]byte), &user)
			ctx.Set("name", user.Name)
			ctx.Set("role", user.Role)
			ctx.Set("age", user.Age)
			ctx.Set("vip", user.Vip)
		}

	}

}

func main() {
	logger := zap.NewExample()

	engine := gin.Default()
	engine.GET("/login", login)
	engine.GET("/postVideo", checkAuth(), postVideo)

	engine.Run(":10101")
	logger.Info("Server started on address localhost:10101")
}
