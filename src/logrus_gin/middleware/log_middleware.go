package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	status200 = 42
	status404 = 43
	status500 = 41
	//methodGET = 44
)

func LogMiddleware() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		// Process request
		c.Next()
		// Stop timer
		TimeStamp := time.Now()
		TimeSub := TimeStamp.Sub(start)
		ClientIP := c.ClientIP()
		Method := c.Request.Method
		StatusCode := c.Writer.Status()
		//BodySize := c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		var statusColor string
		switch StatusCode {
		case 200:
			statusColor = fmt.Sprintf("\033[%dm %d \033[0m", status200, StatusCode)
		case 404:
			statusColor = fmt.Sprintf("\033[%dm %d \033[0m", status404, StatusCode)
		case 500:
			statusColor = fmt.Sprintf("\033[%dm %d \033[0m", status500, StatusCode)
		}
		logrus.Infof("[yf] %s | %s | %v | %v | \033[44m %s \033[0m %s",
			start.Format("2006/01/02 - 15:04:05"),
			statusColor,
			TimeSub,
			ClientIP,
			Method,
			path,
		)
	}
}
