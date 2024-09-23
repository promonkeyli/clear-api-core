package utils

import (
	"encoding/json"
	"example.com/m/v2/src/public"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
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
	}
	// 转json对象
	var openAPI openapi3.T
	err := json.Unmarshal([]byte(config.OpenAPI), &openAPI)
	if err != nil {
		Log("err"+err.Error(), Red)
		return
	}
	// 3. 创建 api文件夹，如果不存在，则生成
	os.Mkdir(path, 0777)
	// 4. 创建一个index.ts总导出文件
	CreateApiMain(openAPI, path)
	// 5. 创建一个typings.d.ts文件
	CreateApiType(openAPI, path)
	// 6. 创建多个接口文件
	CreateApiFiles(openAPI, path, config.RequestLibPath)
}
