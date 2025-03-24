package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	//file, _ := os.OpenFile("./document/logrus.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	file, _ := os.Create("./document/logrus.log") // 是覆盖
	// 传入实现Writer的结构体实例
	writers := []io.Writer{
		file,
		os.Stdout,
	}
	//logrus.SetOutput(file)
	// 同时控制台和日志输出
	logrus.SetOutput(io.MultiWriter(writers...))
	logrus.Infof("信息")
	logrus.Error("错误") // 在日志里也会换行
	logrus.Errorf("错误 %s", "格式化")
	logrus.Errorln("错误 换行")
}

/*func SetOutput(out io.Writer) {
	std.SetOutput(out)
}*/
/*func Create(name string) (*File, error) {
	return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
}*/
