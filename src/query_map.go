package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		// 查询参数
		//func (c *Context) QueryMap(key string) (dicts map[string]string)
		ids := c.QueryMap("ids")
		// 表单参数
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v\n", ids, names)
	})
	router.Run(":8090")
}
