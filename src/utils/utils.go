package utils

import "os"

// 判断文件夹是否存在
func dirExists(path string) bool {
	info, err := os.Stat(path)
	// 如果文件夹不存在或发生其他错误
	if os.IsNotExist(err) {
		return false
	}
	// 确保路径是目录
	return info.IsDir()
}
