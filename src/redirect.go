package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})
	r.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		// 根据gin.context中的请求方法和路径，手动匹配路由并执行对应的处理函数
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})
	r.Run(":8090")
}
