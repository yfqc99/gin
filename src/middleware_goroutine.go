package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// 创建在 goroutine 中使用的副本(上下文的副本)
		//直接使用原始的 gin.Context 可能会导致数据竞争问题。通过使用 c.Copy() 创建副本，可以在协程中安全地使用上下文
		cCp := c.Copy()
		go func() {
			// 用 time.Sleep() 模拟一个长任务。
			time.Sleep(5 * time.Second)

			// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
		//[GIN] 2025/03/24 - 11:24:37 | 200 |            0s |       127.0.0.1 | GET      "/long_async"
		//2025/03/24 11:24:42 Done! in path /long_async
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// 用 time.Sleep() 模拟一个长任务。
		time.Sleep(5 * time.Second)

		// 因为没有使用 goroutine，不需要拷贝上下文
		log.Println("Done! in path " + c.Request.URL.Path)
		//2025/03/24 11:25:05 Done! in path /long_sync
		//[GIN] 2025/03/24 - 11:25:05 | 200 |    5.0146727s |       127.0.0.1 | GET      "/long_sync"
	})

	// 监听并在 0.0.0.0:8090 上启动服务
	r.Run(":8090")
	//在 Web 服务器场景下，main 函数通常会阻塞在服务器的运行上，不会提前结束
}
