package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		// 获取表单中message数据
		message := c.PostForm("message")
		// 获取表单中nick数据，若该字段不存在，则使用默认值anonymous
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8090")
}
