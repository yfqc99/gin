package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}
type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func startPage(c *gin.Context) {
	var user UserInfo
	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。
	if c.ShouldBind(&user) == nil {
		log.Println(user.Name)
		log.Println(user.Age)
		log.Println(user.Sex)
	}

	c.String(200, "Success")
}
func main() {
	// bind可以将前端传递来的数据与结构体进行绑定和参数校验
	//must bind 校验失败会修改状态码
	router := gin.Default()
	/*router.GET("/", func(context *gin.Context) {
		context.BindJSON()
	})
	func (c *Context) BindJSON(obj any) error {
		return c.MustBindWith(obj, binding.JSON)
	}*/
	///testing?name=%E5%BC%A0%E4%B8%89&age=18&sex=%E5%A5%B3
	// 张三
	// 18
	// 女
	router.GET("/testing", startPage)

	// 绑定json数据
	router.POST("/", func(context *gin.Context) {
		var userInfo UserInfo
		err := context.ShouldBindJSON(&userInfo)
		if err != nil {
			context.JSON(404, "错误信息")
			return
		}
		context.JSON(200, userInfo)
		/*{
		"name": "张三",
		"age": 18,
		"sex": "女"
		}*/
	})

	// 绑定查询参数，在字段后添加form的tag
	router.POST("/query", func(context *gin.Context) {
		var userInfo UserInfo
		err := context.ShouldBindQuery(&userInfo)
		if err != nil {
			context.JSON(404, "错误信息")
			return
		}
		context.JSON(200, userInfo)
		// 传入的参数，要可以转换为对应类型，否则无法转换
		//url编码 http://127.0.0.1:8090/query?name=%E6%9D%8E%E5%9B%9B&age=20&sex=%E7%94%B7
		/*{
			"name": "李四",
			"age": 20,
			"sex": "男"
		}*/
		//http://127.0.0.1:8090/query?name=%E6%9D%8E%E5%9B%9B&age=%E5%A5%B3&sex=%E7%94%B7
		//"错误信息"
	})

	// 动态参数 uri
	router.POST("/uri/:name/:age/:sex", func(context *gin.Context) {
		var userInfo UserInfo
		err := context.ShouldBindUri(&userInfo)
		if err != nil {
			context.JSON(404, "错误信息")
			return
		}
		context.JSON(200, userInfo)
		// apifox不能直接填写中文和特殊字符
		// http://127.0.0.1:8090/uri/wu/22/nan
		/*{
			"name": "wu",
			"age": 22,
			"sex": "nan"
		}*/
		// 对其手动进行url编码后填上去
		// http://127.0.0.1:8090/uri/%E9%92%B1%E4%BA%94/22/%E7%94%B7
		/*{
			"name": "钱五",
			"age": 22,
			"sex": "男"
		}*/
	})

	// 绑定form-data，通过form tag
	router.POST("/form", func(context *gin.Context) {
		var userInfo UserInfo
		var user UserInfo
		// 根据请求头类型查找
		// 显式绑定声明绑定，绑定表单数据到userInfo结构体
		//context.ShouldBindWith(&userInfo, binding.Form)
		// 自动绑定
		err := context.ShouldBind(&userInfo)
		if err != nil {
			context.JSON(404, "错误信息")
			return
		}
		// 当类型是JSON, XML, MsgPack, ProtoBuf等不可以重复调用
		// EOF
		// 类型是Query, Form, FormPost, FormMultipart等 可以重复调用
		//userInfo： {周六 18 女}
		//user： {周六 18 女}
		err = context.ShouldBind(&user)
		if err != nil {
			context.JSON(404, err.Error())
			return
		}
		fmt.Println("userInfo：", userInfo)
		fmt.Println("user：", user)

		context.JSON(200, userInfo)
		// form-data
		/*{
			"name": "周六",
			"age": 24,
			"sex": "女"
		}*/
		//x-www-form-urlencoded
		/*{
			"name": "周六",
			"age": 24,
			"sex": "女"
		}*/
		//json
		/*{
			"name": "周六",
			"age": 24,
			"sex": "女"
		}*/
	})

	// 自动识别绑定类型
	router.POST("/bind", func(context *gin.Context) {
		var userInfo UserInfo
		err := context.Bind(&userInfo)
		if err != nil {
			context.JSON(404, "错误信息")
			return
		}
		context.JSON(200, userInfo)
		///bind?name=%E5%BC%A0%E4%B8%89&age=20&sex=%E7%94%B7
		//{"name":"张三","age":20,"sex":"男"}
		// multipart/form-data
		///bind
		/*{
			"name": "周六",
			"age": 18,
			"sex": "女"
		}*/
		//json
		/*{
			"name": "李四",
			"age": 22,
			"sex": "男"
		}*/
	})

	// 注册路由，能够处理所有HTTP方法
	router.Any("/any", func(context *gin.Context) {
		var userInfo UserInfo
		// 会忽略请求体中提交的表单数据
		err := context.ShouldBindQuery(&userInfo)
		if err != nil {
			context.JSON(404, "错误信息")
			return
		}
		context.JSON(200, userInfo)
	})

	// 实现多次调用
	router.POST("/some", func(c *gin.Context) {
		objA := formA{}
		objB := formB{}
		// 读取 c.Request.Body 并将结果存入上下文
		//ShouldBindBodyWith 将请求体的内容存储到上下文中，所以可以多次绑定和读取，不会消耗请求体
		if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
			c.String(http.StatusOK, `the body should be formA`)
			// 这时, 复用存储在上下文中的 body。
		} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
			c.String(http.StatusOK, `the body should be formB JSON`)
			// 可以接受其他格式
		} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
			c.String(http.StatusOK, `the body should be formB XML`)
		}
	})

	router.Run(":8090")
}

/*func (c *Context) Bind(obj any) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.MustBindWith(obj, b)
}*/
