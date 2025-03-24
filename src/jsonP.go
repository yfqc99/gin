package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func main() {
	r := gin.Default()

	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		// 如果查询参数中存在回调，则将回调添加到响应体中
		// 用浏览器请求 localhost:8080/JSONP?callback=x
		// 将输出：x({"foo":"bar"})
		// 自动检查查询参数中是否有 callback 参数
		// 如果有，则将数据包裹在回调函数中返回
		c.JSONP(http.StatusOK, data)
		//x({"foo":"bar"});
	})

	r.GET("/JSONPValid", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		callback := c.Query("callback")
		//减少安全风险，可以对回调函数名称进行验证，确保它只包含字母、数字和下划线
		if callback != "" && validCallback(callback) {
			c.String(http.StatusOK, "%s(%v)", callback, data)
		} else {
			//<script>alert('你好')</script>
			//{"foo":"bar"}
			c.JSON(http.StatusOK, data)
		}
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8090")
}
func validCallback(callback string) bool {
	// 验证回调函数名称是否只包含字母、数字和下划线
	match, _ := regexp.MatchString(`^[a-zA-Z_$][0-9a-zA-Z_$]*$`, callback)
	return match
}
