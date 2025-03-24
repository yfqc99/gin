package main

import "github.com/sirupsen/logrus"

func main() {
	/*
		// 默认使用text形式输出
		// 设置位json格式
		//logrus.SetFormatter(&logrus.JSONFormatter{})
		//{"ip":"192.168.200.254","level":"error","msg":"错误","time":"2025-03-19T19:41:45+08:00","user_id":"21"}
	*/

	// 设置颜色输出
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	/*
		//FullTimestamp:   true 在连接到tty时输出完成的时间戳
		//ERRO[2025-03-19 19:50:38] 错误
		//WARN[2025-03-19 19:50:38] 警告
		//INFO[2025-03-19 19:50:38] 信息
		//DEBU[2025-03-19 19:50:38] 调试
		//INFO[2025-03-19 19:50:38] 打印
	*/
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Error("错误")   // 红色
	logrus.Warnln("警告")  // 黄色
	logrus.Infof("信息")   // 绿色
	logrus.Debugf("调试")  // 白色
	logrus.Println("打印") // 绿色

}
