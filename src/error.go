package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type User struct {
	// 自定义错误信息
	Name string `json:"name" binding:"required" msg:"用户名校验失败""`
	Age  int    `json:"age" binding:"required" msg:"年龄校验失败""`
}

// 通过反射获取到msg标签里的内容
func GetValidMsg(err error, obj interface{}) string {
	getobj := reflect.TypeOf(obj)
	// 把err接口断言为具体类型
	//type ValidationErrors []FieldError
	if errors, ok := err.(validator.ValidationErrors); ok {
		// 断言成功，防止其他类型的错误
		// errors 是切片，因为可能会存在多个错误
		for _, e := range errors {
			// 循环获取每个错误信息
			// e.Field()返回字段名称，标签名称优先于字段的实际名称
			// 根据报错的字段，获取结构体的具体字段
			if f, exist := getobj.Elem().FieldByName(e.Field()); exist {
				// 只获取第一个报错信息
				return f.Tag.Get("msg")
			}
		}
	}
	return err.Error()
}

func main() {
	router := gin.Default()
	router.POST("/", func(context *gin.Context) {
		var user User
		err := context.ShouldBindJSON(&user)
		/*type error interface {
			Error() string
		}*/
		if err != nil {
			//reflect: Elem of invalid type main.User 不传入指针类型会报错
			//context.JSON(404, GetValidMsg(err, user))
			context.JSON(404, GetValidMsg(err, &user))
			return
		}
		context.JSON(200, user)
		// 默认错误信息
		//"Key: 'User.Age' Error:Field validation for 'Age' failed on the 'required' tag"
		// "年龄校验失败"
	})
	router.Run(":8090")
}
