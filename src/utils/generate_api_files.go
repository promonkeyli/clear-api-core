package utils

import (
	"example.com/m/v2/src/public"
	"fmt"
	"os"
)

func GenerateAPIFiles(config public.GenFilesConfig) {
	// 1. 获取当前操作路径,并做路径拼接
	pwd, _ := os.Getwd()
	var path string
	if config.OutDir == "" {
		path = pwd + "/api"
	} else {
		path = pwd + "/" + config.OutDir + "/api"
	}
	// 2. 判断该路径下是否存在 api 文件夹，存在删除再创建，不存在直接创建
	if dirExists(path) {
		// 文件夹存在，则递归删除文件夹及其内容
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println("Error deleting directory:", err)
		}
	} else {
		// 文件夹不存在，则生成，多个前端接口文件，一个类型文件，一个index.ts总导出文件
		CreateApiMain(config.OpenAPI)
	}
}
