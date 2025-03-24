package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// 不同目录下名称相同的模版 ** 表示递归匹配子目录
	router.LoadHTMLGlob("templates/**/*")
	// 不添加define end会报错
	//Error #01: html/template: "posts/index.tmpl" is undefined
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	router.Run(":8090")
}
