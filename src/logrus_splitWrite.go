package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
)

type LogFormatter struct {
	Prefix     string
	TimeFormat string
}

type logFileWriter struct {
	file     *os.File
	logPath  string
	fileDate string
	prefix   string
}

const (
	Black  = 0
	Red    = 1
	Green  = 2
	Yellow = 3
	Blue   = 4
	Purple = 5
	Cyan   = 6
	White  = 7
)

func (lf LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var color int
	switch entry.Level {
	case logrus.ErrorLevel:
		color = Red
	case logrus.InfoLevel:
		color = Green
	case logrus.DebugLevel:
		color = Purple
	case logrus.WarnLevel:
		color = Yellow
	default:
		color = White
	}
	var b *bytes.Buffer
	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}
	formatTime := entry.Time.Format(lf.TimeFormat)
	fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
	/*msg := fmt.Sprintf("[%s] \033[3%dm \t[%s] \033[0m \t[%s] \t%s \t%s\n", lf.Prefix, color, entry.Level, formatTime, entry.Message, fileVal)
	return []byte(msg), nil*/
	fmt.Fprintf(b, "[%s] \033[3%dm \t[%s] \033[0m \t[%s] \t%s \t%s\n", lf.Prefix, color, entry.Level, formatTime, entry.Message, fileVal)
	return b.Bytes(), nil
}
func (lw *logFileWriter) Write(data []byte) (int, error) {
	// 判断是否需要修改日期
	fileDate := time.Now().Local().Format("2006-01-02")
	if lw.fileDate != fileDate {
		lw.file.Close()
		fileName := fmt.Sprintf("%s/%s-%s.log", lw.logPath, lw.prefix, fileDate)
		lw.file, _ = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		lw.fileDate = fileDate
	}
	n, err := lw.file.Write(data)
	return n, err
}
func InitLog(logPath, prefix string) {
	// 格式化时间
	fileDate := time.Now().Format("2006-01-02")
	// 创建目录
	os.MkdirAll(logPath, os.ModePerm)
	fileName := fmt.Sprintf("%s/%s-%s.log", logPath, prefix, fileDate)
	file, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	fileWriter := logFileWriter{file, logPath, fileDate, prefix}
	logrus.SetOutput(io.MultiWriter(os.Stdout, &fileWriter))
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&LogFormatter{prefix, "2006-01-02 15:04:05"})
}
func main() {
	// 按照时间对日志进行分割
	InitLog("./document/logs/", "yf")
	logrus.Error("错误")
	logrus.Warnln("警告")
	logrus.Infof("信息")
	logrus.Debugf("调试")
	logrus.Println("打印")
}
