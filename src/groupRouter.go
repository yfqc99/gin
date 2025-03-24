package main

import "github.com/gin-gonic/gin"

type UserList struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type ArticleInfo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func UserListView(ctx *gin.Context) {
	var userList = []UserList{
		{"张三", 18},
		{"李四", 20},
		{"钱五", 22},
	}
	ctx.JSON(200, userList)
}
func ArticleListView(ctx *gin.Context) {
	var articleList = []ArticleInfo{
		{"python", "文章1"},
		{"java", "文章2"},
		{"go", "文章3"},
	}
	ctx.JSON(200, articleList)
}
func UserRouterInit(api *gin.RouterGroup) {
	// 嵌套
	userManger := api.Group("user_manger")
	{
		userManger.GET("/users2", UserListView)
	}
}
func ArticleRouterInit(api *gin.RouterGroup) {
	articleManger := api.Group("article_manger")
	{
		articleManger.GET("/articles", ArticleListView)
		/*[
			{
			"title": "python",
			"content": "文章1"
			},
			{
			"title": "java",
			"content": "文章2"
			},
			{
			"title": "go",
			"content": "文章3"
			}
		]*/
	}
}
func main() {
	router := gin.Default()
	// 普通操作
	//router.GET("/users", UserListView)
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
	// 对指定组进行统一的操作
	api := router.Group("api")
	{
		//http://127.0.0.1:8090/users 404
		//http://127.0.0.1:8090/api/users 正常响应
		api.GET("/users", UserListView)
		api.POST("/users1", UserListView)
	}
	UserRouterInit(api)
	ArticleRouterInit(api)
	router.Run(":8090")
}
