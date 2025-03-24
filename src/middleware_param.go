package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User2 struct {
	Name string
	Age  int
}

func m6(ctx *gin.Context) {
	fmt.Println("m6")
	// 传递参数
	// 第二个参数是any类型，可以传递任意数据
	//ctx.Set("name", "张三")
	ctx.Set("user", User2{
		"张三",
		18,
	})
}
func m7(ctx *gin.Context) {
	fmt.Println("m7")
}
func index1(ctx *gin.Context) {
	fmt.Println("index")
	//fmt.Println(ctx.Get("name"))
	user, _ := ctx.Get("user")
	fmt.Println(user.(User2).Name)
	fmt.Println(user.(User2).Age)
	ctx.JSON(200, "index1")
	/*m6
	m7
	index
	张三
	18*/
}
func main() {
	router := gin.Default()
	router.Use(m6, m7)
	router.GET("/", index1)
	/*m6
	m7
	index
	张三 true*/
	router.Run(":8090")
}
