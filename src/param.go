package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 查询参数
func _query(ctx *gin.Context) {
	user := ctx.Query("user")
	fmt.Println("user:", user)
	// user:
	// /query
	// 自行进行url解码
	// user: 张三
	// 自动进行url编码
	// /query?user=%E5%BC%A0%E4%B8%89
	// 是否有查询参数输入
	fmt.Println(ctx.GetQuery("user"))
	// false false前面有空字符串
	//张三 true
	// 获取到指定参数的数组
	fmt.Println(ctx.GetQueryArray("user"))
	/*	user: 张三
		张三 true
		[张三 李四] true*/
	fmt.Println(ctx.GetQueryMap("user"))

	// 当没有传递对应参数使用默认值
	page := ctx.DefaultQuery("page", "0")
	fmt.Println(page)
}

// 参数值是动态变化的uri
func _param(ctx *gin.Context) {
	// 任意类型
	fmt.Println(ctx.Param("user_id"), " ", ctx.Param("book_id"))
	//1
	// /param/1
	//1   66a
	// /param/1/66a
}

// 表单 可以接收multipart/form-data和application/x-www-form-urlencoded
func _form(ctx *gin.Context) {
	fmt.Println(ctx.PostForm("name"))
	//张三
	fmt.Println(ctx.PostFormArray("name"))
	/*张三
	[张三 李四]*/
	// 当没有表单没有传入对应参数的值时，会使用默认值
	fmt.Println("age：", ctx.DefaultPostForm("age", "18"))
	// 18
	// 当传入空字符串时不会使用默认值
	// age：
	// 以上不能接收文件形式参数
	// 接收所有form参数，包括文件
	forms, _ := ctx.MultipartForm()
	fmt.Println(forms)
	//&{map[age:[22] name:[张三 李四]] map[file:[0xc000178660]]}
}

// 原始参数
func _raw(ctx *gin.Context) {
	//application/x-www-form-urlencoded 形式
	//fmt.Println(ctx.GetRawData())
	//[110 97 109 101 61 37 69 53 37 66 67 37 65 48 37 69 52 37 66 56 37 56 57 38 97 103 101 61 49 56] <nil>
	//body, _ := ctx.GetRawData()
	// 只能获取一次，重复获取只有第一次有数值
	//fmt.Println("body：", body, " ", "string：", string(body))
	/*body： [110 97 109 101 61 37 69 53 37 66 67 37 65 48 37 69 52 37 66 56 37 56 57 38 97 103 101 61 49 56]
	string： name=%E5%BC%A0%E4%B8%89&age=18*/
	// 当请求向服务器传入参数时，会对中文和特殊字符自动进行url编码
	// 解码
	//unesc, _ := url.QueryUnescape(string(body))
	//fmt.Println("string：", unesc)
	//string： name=张三&age=18

	// multipart/form-data
	/*string： ----------------------------463604394404091591560379
	Content-Disposition: form-data; name="name"

	张三
	----------------------------463604394404091591560379
	Content-Disposition: form-data; name="age"

	18
	----------------------------463604394404091591560379--*/

	//json
	//会将对应的格式一起传入
	/*string： {
		"name":"张三",
			"age":18
	}*/
	//string： {"name":"张三","age":18}
	// 获得请求头
	fmt.Println(ctx.GetHeader("content-Type"))
	// application/json
	var user user
	_bindJson(ctx, user)
	//map[age:18 name:张三]
}

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func _bindJson(ctx *gin.Context, obj any) {
	body, _ := ctx.GetRawData()
	contentType := ctx.GetHeader("content-type")
	switch contentType {
	case "application/json":
		err := json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println("解析失败", err)
		}
		fmt.Println(obj)
	}
}

func main() {
	router := gin.Default()
	router.GET("/query", _query)
	router.GET("/param/:user_id", _param)
	router.GET("/param/:user_id/:book_id", _param)

	// :name 单个路径段
	// *action 是通配符参数，匹配url中剩余的所有路径段
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		//zhangsan is /18/nv
		c.String(http.StatusOK, message)
	})
	router.POST("/form", _form)
	router.POST("/raw", _raw)
	router.Run(":8090")
}
