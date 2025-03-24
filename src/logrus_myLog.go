package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
)

const (
	cBlack  = 0
	cRed    = 1
	cGreen  = 2
	cYellow = 3
	cBlue   = 4
	cPurple = 5
	cCyan   = 6
	cWhite  = 7
)

type MyFormatter struct {
	Prefix     string
	TimeFormat string
}

func (mf MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var color int
	// 获取日志等级
	switch entry.Level {
	case logrus.ErrorLevel:
		color = cRed
	case logrus.InfoLevel:
		color = cGreen
	case logrus.DebugLevel:
		color = cPurple
	case logrus.WarnLevel:
		color = cYellow
	default:
		color = cWhite
	}
	// 缓冲区
	var b *bytes.Buffer
	// 把内容放入该缓冲区
	if entry.Buffer == nil {
		// 初始化
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}
	// 路径
	// 格式化时间
	//formatTime := entry.Time.Format("2006-01-02 15:04:05")
	formatTime := entry.Time.Format(mf.TimeFormat)
	// 文件路径 文件行号
	// path.Base 只获取最后部分的相对路径
	fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
	// 设置格式
	fmt.Fprintf(b, "[%s] \033[3%dm \t[%s] \033[0m \t[%s] \t%s \t%s\n", mf.Prefix, color, entry.Level, formatTime, entry.Message, fileVal)
	return b.Bytes(), nil
}
func main() {
	// 否则调用entry.Caller是空指针，显示行号定位到日志的具体位置
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&MyFormatter{"yf", "2006-01-02 15:04:05"})
	logrus.Error("错误")
	logrus.Warnln("警告")
	logrus.Infof("信息")
	logrus.Debugf("调试")
	logrus.Println("打印")
}
