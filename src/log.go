package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 输出格式
func LogFormatterParams(params gin.LogFormatterParams) string {
	return fmt.Sprintf(
		"[yf] %s |%s%d%s |%s%s%s %s \n",
		params.TimeStamp.Format("2006-01-02 15:04:05"),
		params.StatusCodeColor(), params.StatusCode, params.ResetColor(),
		//params.Method,
		// 添加颜色
		// 匹配颜色，字段，归位
		params.MethodColor(), params.Method, params.ResetColor(),
		params.Path,
	)
	//[yf] 2025-03-19 11:30:56 |200 |GET /index
}
func main() {
	//using env:   export GIN_MODE=release
	//using code:  gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.ReleaseMode) // 去除gin-debug
	/*
		//file, _ := os.Create("./document/gin.log")

		// 默认的写入
		//var DefaultWriter io.Writer = os.Stdout
		// 改变写入的默认位置,，可以填写多个 io.Writer
		gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

		//var DebugPrintRouteFunc func(httpMethod, absolutePath, handlerName string, nuHandlers int)
		// 自定义路由调试
		gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
			// 请求方法
			// 请求路径
			// 处理函数名
			// 处理函数个数
			log.Printf(
				"[yf] %s %s %s %d \n", httpMethod, absolutePath, handlerName, nuHandlers,
			)
			//2025/03/19 11:19:27 [yf] GET /index main.main.func2 3
		}
		router := gin.Default()
	*/
	// 强制log 高亮输出，默认是高亮(颜色)输出的
	//gin.ForceConsoleColor()
	// 关闭高亮，无法控制自定义的高亮
	gin.DisableConsoleColor()
	//[GIN] 2025/03/23 - 09:23:04 | 200 |            0s |       127.0.0.1 | GET      "/index"
	// 自动包括gin.Recovery
	//router := gin.Default()
	router := gin.New()

	// 通过格式化器，定义日志的输出格式
	//router.Use(gin.LoggerWithFormatter(LogFormatterParams), gin.Recovery())
	//Formatter LogFormatter
	//type LogFormatter func(params LogFormatterParams) string

	// 通过配置结构体，全面定义日志记录器的行为，包括日志的格式，输出位置，日志级别等
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: LogFormatterParams,
		//捕获 HTTP 请求处理过程中发生的任何 panic（运行时恐慌），并恢复程序的正常运行，防止服务器因未处理的异常而崩溃
	}), gin.Recovery())
	router.GET("/index", func(context *gin.Context) {})
	// 打印所有的路由
	fmt.Println(router.Routes())
	// 不写默认是8080
	router.Run(":8090")
}

/*func (engine *Engine) Routes() (routes RoutesInfo) {
	for _, tree := range engine.trees {
		routes = iterate("", tree.method, routes, tree.root)
	}
	return routes
}*/
/*type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	HandlerFunc HandlerFunc
}*/
// 在router.GET 等会调用该方法
/*
func (engine *Engine) addRoute(method, path string, handlers HandlersChain) {
	assert1(path[0] == '/', "path must begin with '/'")
	assert1(method != "", "HTTP method can not be empty")
	assert1(len(handlers) > 0, "there must be at least one handler")

	debugPrintRoute(method, path, handlers)

	root := engine.trees.get(method)
	if root == nil {
		root = new(node)
		root.fullPath = "/"
		engine.trees = append(engine.trees, methodTree{method: method, root: root})
	}
	root.addRoute(path, handlers)

	if paramsCount := countParams(path); paramsCount > engine.maxParams {
		engine.maxParams = paramsCount
	}

	if sectionsCount := countSections(path); sectionsCount > engine.maxSections {
		engine.maxSections = sectionsCount
	}
}
func debugPrintRoute(httpMethod, absolutePath string, handlers HandlersChain) {
	if IsDebugging() {
		nuHandlers := len(handlers)
		handlerName := nameOfFunction(handlers.Last())
		if DebugPrintRouteFunc == nil {
			debugPrint("%-6s %-25s --> %s (%d handlers)\n", httpMethod, absolutePath, handlerName, nuHandlers)
		} else {
			DebugPrintRouteFunc(httpMethod, absolutePath, handlerName, nuHandlers)
		}
	}
}
*/
//func LoggerWithConfig(conf LoggerConfig) HandlerFunc {
//	formatter := conf.Formatter
//	if formatter == nil {
//		formatter = defaultLogFormatter
//	}
//
//	out := conf.Output
//	if out == nil {
//		out = DefaultWriter
//	}
//	......
//}
