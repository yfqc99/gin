package main

import "github.com/sirupsen/logrus"

func main() {
	// 设置调试位最低等级
	logrus.SetLevel(logrus.DebugLevel)
	/*
		//time="2025-03-19T17:07:10+08:00" level=error msg="错误"
		//time="2025-03-19T17:07:10+08:00" level=warning msg="警告"
		//time="2025-03-19T17:07:10+08:00" level=info msg="信息"
		//time="2025-03-19T17:07:10+08:00" level=debug msg="调试"
		//time="2025-03-19T17:07:10+08:00" level=info msg="打印"
	*/
	logrus.Error("错误")
	logrus.Warnln("警告")
	logrus.Infof("信息")
	logrus.Debugf("调试")
	logrus.Println("打印")
	/*
		// 默认等级，比info小的不会输出
		//time="2025-03-19T17:03:50+08:00" level=error msg="错误"
		//time="2025-03-19T17:03:50+08:00" level=warning msg="警告"
		//time="2025-03-19T17:03:50+08:00" level=info msg="信息"
		//time="2025-03-19T17:03:50+08:00" level=info msg="打印"
	*/
	// 获取最小等级
	logrus.GetLevel()
}
