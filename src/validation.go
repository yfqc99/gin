package main

import "github.com/gin-gonic/gin"

type SignUserInfo struct {
	//Name       string `json:"name" binding:"required"`// 必填项
	Name       string `json:"name" binding:"min=4,max=6"` //最小4个字符，最大六个字符
	Age        int    `json:"age" binding:"lt=30,gt=18"`  // 大于18，小于30
	Password   string `json:"password"`
	RePassword string `json:"re_password" binding:"eqfield=Password"` // 校验两个指定字段是否相同，没有为空校验
	// 要求RePassword必须大于Password
	//RePassword string `json:"re_password" binding:"gtfield=Password"`
}

func main() {
	router := gin.Default()
	router.POST("/", func(context *gin.Context) {
		var user SignUserInfo
		err := context.ShouldBindJSON(&user)
		if err != nil {
			context.JSON(404, gin.H{"msg": err.Error()})
			return
		}
		context.JSON(200, user)
		// required
		/*{
			"name": "张三",
			"age": 18,
			"password": "123",
			"re_password": "123"
		}*/
		// 不传入name字段或者值为空
		/*[
			{}
		]*/
		// max和min，utf-8中文3个字节
		// 李，7个字符不匹配 李123456
		/*{
			"msg": [
				{}
			]
		}*/
		// 4个字符匹配
		/*{
			"name": "李123",
			"age": 18,
			"password": "123",
			"re_password": "123"
		}*/
		// lt和gt 不包含两端
		// lte和gte 包含两端
		/*{
			"msg": [
				{}
			]
		}*/
		/*{
			"name": "李123",
			"age": 19,
			"password": "123",
			"re_password": "123"
		}*/
		//eqfield 指定字段是否相同
		/*{
			"name": "李123",
			"age": 19,
			"password": "123",
			"re_password": "123"
		}*/
		// 123 和 1
		/*{
			"msg": [
				{}
			]
		}*/
		/*{
			"name": "李123",
			"age": 19,
			"password": "",
			"re_password": ""
		}*/
	})
	router.Run(":8090")
}
