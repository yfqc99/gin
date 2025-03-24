package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	allLog  = "all"
	errLog  = "err"
	warnLog = "warn"
	infoLog = "info"
)

type FileLevelHook struct {
	file     *os.File
	errFile  *os.File
	warnFile *os.File
	infoFile *os.File
	logPath  string
}

func (FileLevelHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (fh FileLevelHook) Fire(entry *logrus.Entry) error {
	msg, _ := entry.String()
	switch entry.Level {
	case logrus.InfoLevel:
		fh.infoFile.Write([]byte(msg))
	case logrus.ErrorLevel:
		fh.errFile.Write([]byte(msg))
	case logrus.WarnLevel:
		fh.warnFile.Write([]byte(msg))
	default:
		fh.file.Write([]byte(msg))
	}
	return nil
}
func initLevel(logPath string) {
	os.MkdirAll(logPath, os.ModePerm)
	infoFile, _ := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, infoLog), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	errFile, _ := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, errLog), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	warnFile, _ := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, warnLog), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	allFile, _ := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, allLog), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	fileHook := FileLevelHook{
		allFile,
		errFile,
		warnFile,
		infoFile,
		logPath,
	}
	logrus.AddHook(fileHook)
	logrus.SetLevel(logrus.DebugLevel)
}
func main() {
	initLevel("./document/logLevels")
	logrus.Error("错误")
	logrus.Warnln("警告")
	logrus.Infof("信息")
	logrus.Debugf("调试")
	logrus.Println("打印")
}
