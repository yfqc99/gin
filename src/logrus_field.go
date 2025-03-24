package main

import "github.com/sirupsen/logrus"

func server() {

}

func main() {
	// 设置通用字段
	//log := logrus.WithField("app", "study").WithField("service", "logrus")
	//time="2025-03-19T17:16:00+08:00" level=error msg="错误" app=study service=logrus
	//type Fields map[string]interface{}
	log := logrus.WithFields(logrus.Fields{
		"user_id": "21",
		"ip":      "192.168.200.254",
	})
	/*time="2025-03-19T17:21:45+08:00" level=error msg="错误" ip=192.168.200.254 user_id=21*/
	// 添加通用字段
	log.Errorf("错误")
	// 在处理http请求时，上下文中，所有的日志都会有request_id 和 user_id
}
