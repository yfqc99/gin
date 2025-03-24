package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type FileDataHook struct {
	file     *os.File
	logPath  string
	fileDate string
	prefix   string
}

type MyFormatter struct{}

func (mf MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	// 把内容放入该缓冲区
	if entry.Buffer == nil {
		// 初始化
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}
	// 设置格式
	fmt.Fprintf(b, "%s\n", entry.Message)
	return b.Bytes(), nil
}

func (FileDataHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (fh FileDataHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02_15-04")
	if timer != fh.fileDate {
		// 关闭之前的文件
		fh.file.Close()
		// 创建新文件
		fh.file, _ = os.OpenFile(fmt.Sprintf("%s/%s-%s.log", fh.logPath, fh.prefix, fh.fileDate), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		fh.fileDate = timer
	}
	msg, _ := entry.String()
	fh.file.Write([]byte(msg))
	return nil
}

func InitFile(logPath, prefix string) {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&MyFormatter{})
	fileDate := time.Now().Format("2006-01-02_15-04")
	// 创建目录
	os.MkdirAll(logPath, os.ModePerm)
	fileName := fmt.Sprintf("%s/%s-%s.log", logPath, prefix, fileDate)
	file, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	fileHook := FileDataHook{
		file,
		logPath,
		fileDate,
		prefix,
	}
	logrus.AddHook(fileHook)
}
