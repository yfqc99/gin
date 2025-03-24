package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func m4(ctx *gin.Context) {
	fmt.Println("m4...request")
	//fmt.Println("m4...")
	ctx.Next()
	fmt.Println("m4...response")
}
func m5(ctx *gin.Context) {
	fmt.Println("m5...request")
	//fmt.Println("m5...")
	ctx.Next()
	fmt.Println("m5...response")
}
func main() {
	router := gin.Default()
	// 定义全局中间件
	router.Use(m4, m5)
	// 404 也会执行
	//m4...
	router.GET("/m4", func(context *gin.Context) {
		//fmt.Println("m4")
		fmt.Println("index...request")
		context.JSON(200, "m4")
		context.Next()
		fmt.Println("index...response")
		/*m4...request
		m5...request
		index...request
		index...response
		m5...response
		m4...response*/
	})
	/*m4...
	m4*/
	router.GET("/m5", func(context *gin.Context) {
		fmt.Println("m5")
		context.JSON(200, "m5")
	})
	/*m4...
	m5*/
	// 无用
	//router.Use(m4)
	router.Run(":8090")
}
