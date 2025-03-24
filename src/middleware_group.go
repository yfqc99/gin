package main

import (
	"github.com/gin-gonic/gin"
)

type UserList1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func UserListView1(ctx *gin.Context) {
	var userList = []UserList1{
		{"张三", 18},
		{"李四", 20},
		{"钱五", 22},
	}
	ctx.JSON(200, userList)
}
func UserRouterInit1(api *gin.RouterGroup) {
	// 嵌套
	//userManger := api.Group("user_manger").Use(middle)
	// 传入函数返回值需要满足handlerFunc
	userManger := api.Group("user_manger").Use(test("用户校验失败"))
	{
		userManger.GET("/users", UserListView1)
	}
}

// "权限验证失败"
// 设置token为123
/*[
	{
	"name": "张三",
	"age": 18
	},
	{
	"name": "李四",
	"age": 20
	},
	{
	"name": "钱五",
	"age": 22
	}
]*/
func middle(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token == "123" {
		ctx.Next()
		return
	}
	ctx.JSON(200, "权限验证失败")
	ctx.Abort()
}

func test(msg string) func(ctx *gin.Context) {
	// 当调用该函数做中间件，会被立刻执行
	// 装饰器
	//name := "周六" // 形成一个闭包
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "123" {
			ctx.Next()
			return
		}
		ctx.JSON(200, msg)
		ctx.Abort()
	}
}
func main() {
	//router := gin.Default()
	router := gin.New()
	//参数是 type LogFormatter func(params LogFormatterParams) string
	// 自定义日志
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return ""
	}), gin.Recovery())
	api := router.Group("api")
	// 不经过权限验证
	api.GET("/login", func(context *gin.Context) {
		panic("手动报错")
		context.JSON(200, "登录")
	})
	//"登录"
	UserRouterInit1(api)
	router.Run(":8090")
}

/*func Default(opts ...OptionFunc) *Engine {
	debugPrintWARNINGDefault()
// 不含任何中间件的路由
	engine := New()
// Logger 日志，在控制台打印当前的对应信息
// Recovery 处理panic报错，响应500，但不会影响其他的请求
	engine.Use(Logger(), Recovery())
	return engine.With(opts...)
}*/
