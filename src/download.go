package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/download", func(context *gin.Context) {
		// 可能是需要清理浏览器缓存的，状态码是304
		// 文件没有被下载，而是在浏览器进行展示(打开)
		//context.File("./document/图解HTTP_(上野宣)_(Z-Library).pdf") // 没有返回值
		// 表示文件流，唤醒浏览器下载
		context.Header("content-type", "application/octet-stream")
		// 下载文件，默认文件名是download
		//context.File("./document/图解HTTP_(上野宣)_(Z-Library).pdf")
		// 设置文件名，指定下载的文件名
		context.Header("content-disposition", "attachment;filename="+"测试文件.pdf")
		// 传输过程中的编码形式，可能会造成乱码问题
		//context.Header("content-transfer-encoding", "binary")
		context.File("./document/图解HTTP_(上野宣)_(Z-Library).pdf")
	})

	router.Run(":8090")
}
