package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 开启并监听服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//服务器启动失败且错误不是 http.ErrServerClosed（服务器正常关闭时的错误），则记录致命错误并退出
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	// 接收操作系统信号
	quit := make(chan os.Signal)
	// 捕获os.Interrupt信号，注册要接收到的信号，在这个方法里面会开启一个协程监听信号
	//watchSignalLoopOnce.Do(func() {
	//	if watchSignalLoop != nil {
	//		go watchSignalLoop() // 阻塞等待信号，接收到后根据注册通道和信号映射关系，将信号分发到相应通道中
	//	}
	//})
	signal.Notify(quit, os.Interrupt)
	// 阻塞等待信号
	<-quit
	log.Println("Shutdown Server ...")

	// 当服务器发出关闭信号后在指定的超时时间内，没有处理完所有的请求，服务器会强制关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//srv.Shutdown：使用上下文关闭服务器，允许正在处理的请求完成
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	// 发送中断信号后处理在指定时间内处理完请求关闭
	//[GIN] 2025/03/24 - 14:42:24 | 200 |    5.0008747s |       127.0.0.1 | GET      "/"
	//2025/03/24 14:42:32 Shutdown Server ...
	//[GIN] 2025/03/24 - 14:42:36 | 200 |    5.0126315s |       127.0.0.1 | GET      "/"
	//2025/03/24 14:42:36 Server exiting
}
