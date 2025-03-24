package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func m1(ctx *gin.Context) {
	fmt.Println("m1")
	ctx.JSON(200, "m1响应")
	// 拦截，之后的handlerFunc不再执行
	//ctx.Abort()
}

func m2(ctx *gin.Context) {
	fmt.Println("m2...request")
	ctx.Next()
	fmt.Println("m2...response")
	ctx.JSON(200, "m2响应")
}

// 在响应中间件中添加拦截，并没有用
func m3(ctx *gin.Context) {
	// 请求中间件
	fmt.Println("m3...request")
	ctx.Next()
	// 响应中间件
	ctx.Abort()
	fmt.Println("m3...response")
}

/*
m3...request
index...request
m2...request
m2...response
index...response
m3...response
*/
func index(ctx *gin.Context) {
	fmt.Println("index...request")
	ctx.JSON(200, "测试")
	//ctx.Abort()
	ctx.Next()
	fmt.Println("index...response")
}

// 在请求中间件部分添加ctx.Abort()会拦截之后的handlerFunc不执行
/*
m3...request
index...request
index...response
m3...response
*/
func main() {
	// 中间件
	// 在处理请求时，加入自己的钩子(hook)函数
	// 中间件适合处理一些公共的业务逻辑
	router := gin.Default()
	// 可以传入多个HandlerFunc
	/*router.GET("/", m3, index, func(context *gin.Context) {
		fmt.Println(1)
	}, func(context *gin.Context) {
		fmt.Println(2)
	}, func(context *gin.Context) {
		fmt.Println(3)
	}, m2)*/
	// 每个都被依次执行，可以理解为每个都是一个中间件，都可以响应
	// type HandlerFunc func(*Context) 必须是该类型
	//m1
	//index
	//1
	//2
	//3

	router.GET("/", m3, index, m2)
	// 响应
	//"m1响应""测试"
	//"m1响应""测试""m2响应"
	/*m3...request
	index...request
	m2...request
	m2...response
	index...response
	m3...response*/

	router.Run(":8090")
}

/*func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.handle(http.MethodGet, relativePath, handlers)
}*/
