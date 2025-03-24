package main

import "github.com/gin-gonic/gin"

type SignUserInfo2 struct {
	//Name string `json:"name" binding:"contains=f"` // 包含
	//Name string `json:"name" binding:"excludes=f"` // 不包含
	//Name string `json:"name" binding:"startswith=fg"` // 前缀
	Name     string   `json:"name" binding:"endswith=fg"` // 后缀
	Age      int      `json:"age"`
	Sex      string   `json:"sex" binding:"oneof=man woman"`                     //枚举
	LikeList []string `json:"like_list" binding:"required,dive,startswith=like"` // 对每个值进行校验
	IP       string   `json:"ip" binding:"ip"`
	Url      string   `json:"url" binding:"url"`
	Uri      string   `json:"uri" binding:"uri"`
	Date     string   `json:"date" binding:"datetime=2006-01-02 15:04:05"`
}

func main() {
	router := gin.Default()
	router.POST("/", func(context *gin.Context) {
		var user SignUserInfo2
		err := context.ShouldBindJSON(&user)
		if err != nil {
			context.JSON(404, gin.H{"msg": err.Error()})
			return
		}
		context.JSON(200, user)
		// oneof枚举
		/*{
			"name": "李123",
			"age": 19,
			"sex": "man"
		}
		{
			"name": "李123",
			"age": 19,
			"sex": "woman"
		}*/
		// contains 必须包含指定字段
		/*{
			"name": "李f",
			"age": 19,
			"sex": "woman"
		}*/
		//excludes 不包含指定字段
		/*{
			"name": "李",
			"age": 19,
			"sex": "woman"
		}*/
		//startswith 包含指定前缀
		/*{
			"name": "fg李",
			"age": 19,
			"sex": "woman"
		}*/
		//endswith 包含指定后缀
		/*{
			"name": "李fg",
			"age": 19,
			"sex": "woman"
		}*/
		//dive 对数组中每个值进行后面指定操作的校验
		// 必须前缀是like
		/*	{
			"name": "李fg",
			"age": 19,
			"sex": "woman",
			"like_list": [
				"like_pingpong",
				"like_football"
			]
		}*/
		// ip 该字段填写合法的IP地址，且不能为空
		/*	{
			"name": "李fg",
			"age": 19,
			"sex": "woman",
			"like_list": [
				"like_pingpong",
				"like_football"
			],
			"ip": "127.0.0.1"
		}*/
		// url 该字段填写合法的url，且不能为空
		/*	{
			"name": "李fg",
			"age": 19,
			"sex": "woman",
			"like_list": [
				"like_pingpong",
				"like_football"
			],
			"ip": "127.0.0.1",
			"url": "https://localhost:8090"
		}*/
		// uri 该字段填写合法的uri，且不能为空
		/*{
			"name": "李fg",
			"age": 19,
			"sex": "woman",
			"like_list": [
				"like_pingpong",
				"like_football"
			],
			"ip": "127.0.0.1",
			"url": "https://localhost:8090/123/456/789",
			"uri": "https://localhost:8090/123/456/789"
		}*/
		/*{
			"name": "李fg",
			"age": 19,
			"sex": "woman",
			"like_list": [
				"like_pingpong",
				"like_football"
			],
			"ip": "127.0.0.1",
			"url": "https://localhost:8090/123/456/789",
			"uri": "/123/456/789"
		}*/
		// datetime 输入指定格式的时间，不能为空，可以不合法但位数要匹配
		// 2024-04-18 13:14:520 不可以
		/*{
			"name": "李fg",
			"age": 19,
			"sex": "woman",
			"like_list": [
				"like_pingpong",
				"like_football"
			],
			"ip": "127.0.0.1",
			"url": "https://localhost:8090/123/456/789",
			"uri": "/123/456/789",
			"date": "2024-04-18 13:14:52"
		}*/
		/*{
			"name": "李fg",
			"age": 19,
			"sex": "woman",
			"like_list": [
				"like_pingpong",
				"like_football"
			],
			"ip": "127.0.0.1",
			"url": "https://localhost:8090/123/456/789",
			"uri": "/123/456/789",
			"date": "2024-04-18 13:14:70"
		}*/
	})
	router.Run(":8090")
}
