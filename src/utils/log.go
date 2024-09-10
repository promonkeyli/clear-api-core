package utils

import "fmt"

type Color string

// 使用 iota 定义 ANSI 颜色代码的枚举常量
const (
	Reset  Color = "\033[0m"
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Blue   Color = "\033[34m"
	Purple Color = "\033[35m"
	Cyan   Color = "\033[36m"
	White  Color = "\033[37m"
)

func Log(text string, color Color) {
	if color == "" {
		// 设定color颜色默认值为绿色
		color = Green
	}
	var str = string(color) + ("[API]: ") + text + string(Reset)
	fmt.Println(str)
}
