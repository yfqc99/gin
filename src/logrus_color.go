package main

import "fmt"

/*
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

	func PrintColor(colorCode int, text string, isBackGround bool) {
		if isBackGround {
			fmt.Printf("\033[4%dm %s \033[0m\n", colorCode, text)
		} else {
			fmt.Printf("\033[3%dm %s \033[0m\n", colorCode, text)
		}
	}
*/
func main() {
	// 使用控制字符，设置字体颜色
	fmt.Printf("\033[31m %s \033[0m\n", "红色")
	fmt.Printf("\033[32m %s \033[0m\n", "绿色")
	fmt.Printf("\033[33m %s \033[0m\n", "黄色")
	fmt.Printf("\033[34m %s \033[0m\n", "蓝色")
	fmt.Printf("\033[35m %s \033[0m\n", "紫色")
	fmt.Printf("\033[36m %s \033[0m\n", "青色")
	fmt.Printf("\033[37m %s \033[0m\n", "白色")
	fmt.Printf("\033[30m %s \033[0m\n", "黑色")

	// 背景颜色
	fmt.Printf("\033[41m %s \033[0m\n", "红色")
	fmt.Printf("\033[42m %s \033[0m\n", "绿色")
	fmt.Printf("\033[43m %s \033[0m\n", "黄色")
	fmt.Printf("\033[44m %s \033[0m\n", "蓝色")
	fmt.Printf("\033[45m %s \033[0m\n", "紫色")
	fmt.Printf("\033[46m %s \033[0m\n", "青色")
	fmt.Printf("\033[47m %s \033[0m\n", "白色")
	fmt.Printf("\033[40m %s \033[0m\n", "黑色")
}
