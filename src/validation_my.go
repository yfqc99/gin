package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 验证

type User1 struct {
	Name string `json:"name" binding:"required,sign" msg:"用户名校验失败""`
	Age  int    `json:"age" binding:"required" msg:"年龄校验失败""`
}

func SignValid(fl validator.FieldLevel) bool {
	// 值不能在数组中存在
	nameList := []string{"张三", "李四", "钱五"}
	for _, nameStr := range nameList {
		//func (FieldLevel) Field() reflect.Value
		if nameStr == fl.Field().Interface().(string) {
			return false
		}
	}
	return true
}

func main() {
	router := gin.Default()
	//Engine()注射器
	//func (StructValidator) Engine() any
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册
		// 参数一：名字
		// 参数二：处理函数
		v.RegisterValidation("sign", SignValid)
	}
	router.POST("/", func(context *gin.Context) {
		var user User1
		err := context.ShouldBindJSON(&user)
		if err != nil {
			context.JSON(404, err.Error())
			return
		}
		context.JSON(200, user)
		//"Key: 'User1.Name' Error:Field validation for 'Name' failed on the 'sign' tag"
		/*{
			"name": "周六",
			"age": 18
		}*/
	})
	router.Run(":8090")
}
