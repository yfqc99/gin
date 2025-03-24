package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 格式化日期
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	router := gin.Default()
	// 自定义模版分隔符，用来标识模版中动态内容的边界符号
	// 告诉模板引擎哪部分是需要被解析和替换的变量或逻辑，哪部分是静态内容
	router.Delims("{[{", "}]}")

	// 设置自定义模版函数
	//type FuncMap map[string]any
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("./templates/raw.tmpl")

	router.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2024, 04, 18, 0, 0, 0, 0, time.UTC),
		})
	})
	//Date: 2024/04/18
	router.Run(":8090")
}
