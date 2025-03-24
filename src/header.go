package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// 请求头参数获取
// 响应头
func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		// 不区分大小写
		fmt.Println("1: ", context.GetHeader("user-agent"))
		//1:  Apifox/1.0.0 (https://apifox.com)

		fmt.Println("2: ", context.Request.Header.Get("user-agent"))
		fmt.Println("5: ", context.Request.Header.Get("User-Agent"))
		/*2:  Apifox/1.0.0 (https://apifox.com)
		5:  Apifox/1.0.0 (https://apifox.com)*/

		//type Header map[string][]string
		// 首字母必须大写
		fmt.Println("3: ", context.Request.Header["user-agent"])
		fmt.Println("4: ", context.Request.Header["User-Agent"])
		/*3:  []
		4:  [Apifox/1.0.0 (https://apifox.com)]*/

		// 当一个参数名有多个值时，只能通过直接访问Header，进行全部获取，其他只能获取到第一个值
		/*1:  zhangsan
		2:  zhangsan
		5:  zhangsan
		3:  []
		4:  [zhangsan zhangsan123]*/
		context.JSON(200, "获取成功")
	})

	// 区分爬虫和普通用户
	router.GET("/index", func(context *gin.Context) {
		userAgent := context.GetHeader("user-agent")
		// 正则匹配
		// 字符串包含匹配
		if strings.Contains(userAgent, "python") {
			// 爬虫
			context.JSON(200, "爬虫")
			return
			// python3-requests3.5.8
			//"爬虫"
		}
		context.JSON(200, "普通用户")
		// 在浏览器打开
		// Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Edg/134.0.0.0
		//"普通用户"
	})

	// 响应头
	router.GET("/res", func(context *gin.Context) {
		context.Header("myHeader", "myValue")
		//Myheader	myValue
		context.Header("content-type", "application-text;charset=uft-8")
		//Content-Type	application-text;charset=uft-8
		// 并没有下载，应该是修复了
		//context.JSON(200, "响应头")
		context.JSON(200, gin.H{"data": "响应头"})
	})
	router.Run(":8090")
}
