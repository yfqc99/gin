package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type person struct {
	Name     string `json:"name"`
	Age      int
	Password string `json:"-"`
}

func _string(context *gin.Context) {
	// 状态码不会影响到浏览器的输出
	//context.String(http.StatusNotFound, "respond")
	//context.String(http.StatusOK, "respond")
	// 通过以下方式响应为json字符串
	context.Header("content-type", "application/json")
	context.String(200, `{"age":18,"name":"周六"}`)
	/*	{
		"age": 18,
		"name": "周六"
	}*/
}
func _json(context *gin.Context) {
	// 第二个参数是任意类型的，对该参数进行json转换
	// 注意字段大写!!!!!!!!!!
	/*context.JSON(200, person{
		Age:      18,
		Name:     "张三",
		Password: "123456",
	})*/
	/*	{
		"name": "张三",
		"Age": 18
	}*/

	/*user := map[string]string{
		"age":  "18",
		"name": "钱五",
	}
	context.JSON(200, user)*/
	/*{
		"age": "18",
		"name": "钱五"
	}*/

	// 直接响应json
	// type H map[string]any (any->interface{})
	context.JSON(200, gin.H{
		"age":  22,
		"name": "宋七",
	})
	/*{
		"age": 22,
		"name": "宋七"
	}*/
}

/*
	func (c *Context) JSON(code int, obj any) {
		c.Render(code, render.JSON{Data: obj})
	}
*/
func _xml(context *gin.Context) {
	context.XML(200, gin.H{
		"user":    "李四",
		"message": "hello",
		"status":  http.StatusOK,
	})
	/*	<map>
		<user>李四</user>
		<message>hello</message>
		<status>200</status>
	</map>*/
}
func _yaml(context *gin.Context) {
	context.YAML(200, gin.H{
		"user":    "李四",
		"message": "hello",
		"status":  http.StatusOK,
	})
	/*data:
		message: hello
	  status: 200
	  user: 李四*/
	//浏览器会下载
}
func _html(context *gin.Context) {
	//context.HTML(200, "index.html", nil)
	//context.HTML(200, "index.html", "张三")
	//hello,张三
	/*context.HTML(200, "index.html", gin.H{
		"username": "李四",
	})*/
	//hi,李四
	context.HTML(200, "index.html", person{
		Name:     "钱五",
		Age:      22,
		Password: "123",
	})
	//username：钱五 age：22
}
func _redirect(context *gin.Context) {
	// 重定向状态码301 302 304
	// 301 永久重定向
	// 302 临时重定向
	context.Redirect(302, "https://www.baidu.com")
}
func main() {
	// http.Statusxxx 状态码
	router := gin.Default()
	// 加载指定目录下的所有文档
	router.LoadHTMLGlob("templates/*")

	//router.LoadHTMLFiles("templates/index.html")

	// 在golang中，没有相对文件路径，只有相对项目路径
	// 将具体文件绑定到具体的路由路径上，当访问该路径时，会直接返回指定文件
	// 参数一：网站路径url
	// 参数二：文件路径
	// 访问静态资源
	//http://localhost:8090/png
	router.StaticFile("/png", "static/aaa.png")
	// 可以访问指定存储路径下的所有静态资源
	// 在浏览器页面中，有对应连接，点击查看对应资源
	//http://localhost:8090/static/test.txt
	//http://localhost:8090/static/aaa.png
	// 嵌套层级也可以访问到
	//http://localhost:8090/static/ 该路径下点击test/链接
	//http://localhost:8090/static/test/ 该路径下点击aaa.txt链接
	//http://localhost:8090/static/test/aaa.txt
	// 参数一：网页请求静态目录的前缀
	// 参数二：静态资源目录
	router.StaticFS("/static", http.Dir("static/"))
	router.GET("/", _string)
	router.GET("/json", _json)
	router.GET("/xml", _xml)
	router.GET("/yaml", _yaml)
	router.GET("/html", _html)
	router.GET("/redirect", _redirect)
	router.Run(":8090")
}
