package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	// 创建默认路由，通过操作不同放在不同的路由组中
	router := gin.Default()
	// 绑定路由规则和路由函数，访问/的路由，将由对应函数处理
	router.GET("/", func(context *gin.Context) {
		// 把指定字符串写在响应体中
		context.String(200, "hello")
	})
	// 启动监听，框架把web服务运行在本机的0.0.0.0：8090端口上，本机所有IP
	router.Run(":8090")
	// 源码 就是http.ListenAndServe的进一步封装
	/*func (engine *Engine) Run(addr ...string) (err error) {
		defer func() { debugPrintError(err) }()

		if engine.isUnsafeTrustedProxies() {
			debugPrint("[WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.\n" +
				"Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.")
		}

		address := resolveAddress(addr)
		debugPrint("Listening and serving HTTP on %s\n", address)
		err = http.ListenAndServe(address, engine.Handler())
		return
	}*/
	// 或者 默认的http
	// http.ListenAndServe(":8090",router)
	// 自定义http配置
	s := &http.Server{
		Addr:    ":8090",
		Handler: router,
		// 读和写超时
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// 1MB ，将1的二进制左移20位
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
