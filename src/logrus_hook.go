package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type MyHook struct {
	Writer io.Writer
}

func (MyHook) Levels() []logrus.Level {
	// 所有level都会作用于该hook
	//return logrus.AllLevels
	// 只有error 会作用于该hook
	return []logrus.Level{logrus.ErrorLevel}
}
func (mh MyHook) Fire(entry *logrus.Entry) error {
	//fmt.Println(entry)
	// 在日志中都添加一个字段
	entry.Data["app"] = "yf"
	//time="2025-03-19T21:14:55+08:00" level=error msg="错误" app=yf

	// 将error级别的日志独立输出到error.log文件中
	//entry.Message 只有内容 msg的值
	msg, _ := entry.String()
	mh.Writer.Write([]byte(msg))
	//mh.Writer.Write([]byte(entry.Message))
	return nil
}

// &{0xc000108000 map[] 2025-03-19 21:12:28.1934223 +0800 CST m=+0.029988701 info <nil> 信息 <nil> <nil> }
func main() {
	file, _ := os.OpenFile("./document/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// 实现各种扩展功能
	logrus.AddHook(MyHook{file})
	logrus.Error("错误")
	logrus.Warnln("警告")
	logrus.Infof("信息")
	logrus.Debugf("调试")
	logrus.Println("打印")
}

/*type Hook interface {
	Levels() []Level
	Fire(*Entry) error
}*/
