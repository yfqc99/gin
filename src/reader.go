package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/someDataFromReader", func(c *gin.Context) {
		// 向指定请求路径发送get请求，并获得响应
		response, err := http.Get("https://www.baidu.com")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		//Body io.ReadCloser 获得响应体
		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		// 设置额外响应头
		extraHeaders := map[string]string{
			// 将响应内容作为文件下载，并建议文件名为gopher.png
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		// 从请求体中读取数据，并作为响应返回
		// 适用于处理大文件或者流式数据
		// 直接从请求体中获取数据，不需要将数据加载到内存中
		// 参数一：响应状态码
		// 参数二：读取的数据大小(单位字节)
		// io.Reader 实例，用于从请求体中读取数据
		// 将数据从reader中返回给客户端
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	router.Run(":8090")
}
