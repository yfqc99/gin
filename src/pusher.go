package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

// 创建一个新的模板将指定内容解析到模板中
var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()
	// 将目录下的所有文件挂载到一个路由路径下，当gin会从指定目录中查找并返回相应文件
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		// 获取推送器，并推送/assets/app.js
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用 pusher.Push() 做服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	//HTTPS 服务器，监听端口 8080，并指定 SSL 证书和密钥文件
	// 监听并在 https://127.0.0.1:8090 上启动服务
	r.RunTLS(":8090", "./testdata/server.pem", "./testdata/server.key")
}
