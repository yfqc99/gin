package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求处理之前，状态码默认为200
		t := time.Now()

		// 设置 example 变量
		c.Set("example", "12345")

		// 请求前

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// c.Writer 用于向客户端发送响应
		// 获取发送的 status
		// 获取当前HTTP响应的状态码
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		// 用户获取上下文中指定的值，若指定的键不存在会引发Panic
		// 常用于中间件或处理函数中获取之前设置的值，确保这些值一定存在
		example := c.MustGet("example").(string)

		// 打印："12345"
		log.Println(example)
	})
	/*2025/03/24 10:45:52 12345
	2025/03/24 10:45:52 21.4914ms
	2025/03/24 10:45:52 200*/
	// 监听并在 0.0.0.0:8090 上启动服务
	r.Run(":8090")
}
