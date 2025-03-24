package main

import (
	"gin/src/logrus_gin/log"
	"gin/src/logrus_gin/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	log.InitFile("./document/logGins", "yf")
	router := gin.New()
	router.Use(middleware.LogMiddleware())
	router.GET("/", func(context *gin.Context) {
		logrus.Info("响应成功")
		context.JSON(200, "测试")
	})
	router.Run(":8090")
}
