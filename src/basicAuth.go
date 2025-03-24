package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	r := gin.Default()

	// 使用 Basic Authentication 中间件来保护特定的路由组

	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		// 定义用户名和密码的映射
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	// 需要对用户名和password(username:password)进行base64编码
	// 需要在请求头中添加Authorization，值为Basic username:password
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		//获取通过 BasicAuth 中间件认证的用户名
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			//根据用户名从 secrets 映射中检索私人数据，并返回 JSON 格式的响应
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	// 不添加
	//[GIN] 2025/03/24 - 20:58:09 | 401 |         555µs |       127.0.0.1 | GET      "/admin/secrets"
	// 添加Authorization	Basic Zm9vOmJhcg==
	//[GIN] 2025/03/24 - 21:32:48 | 200 |         999µs |       127.0.0.1 | GET      "/admin/secrets"
	//{
	//	"secret": {
	//	"email": "foo@bar.com",
	//		"phone": "123433"
	//},
	//	"user": "foo"
	//}
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8090")
}
