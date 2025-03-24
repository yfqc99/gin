package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	router := gin.Default()
	// template.Must 处理模版解析中可能出现的错误，遇到错误触发panic
	// template.ParseFiles 解析多个模版文件，并合并到一个模版实例中
	html := template.Must(template.ParseFiles(
		"templates/myTemplate/index.html",
		"templates/myTemplate/footer.html",
		"templates/myTemplate/header.html",
	))
	/*
		<!DOCTYPE html>
		<html>
		<head>
		<title>Main Page</title>
		</head>
		<body>
		<h1>Welcome to our website</h1>

		<p>This is the main content of the page.</p>
		</body>
		</html>
	*/
	// 设置自定义模版渲染器，将解析后的模版设置到路由器中
	router.SetHTMLTemplate(html)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main Page",
		})
	})
	router.Run(":8090")
}
