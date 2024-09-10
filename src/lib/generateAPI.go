package lib

import (
	"example.com/m/v2/src/public"
	"example.com/m/v2/src/utils"
)

func GenerateAPI(config public.ApiConfig) {
	// 1. 根据输入的url 获取swagger json文档
	swagger := utils.GetSwaggerDoc(config.DocUrl)
	// 2. 统一swagger文档至最新的openAPI3.0版本
	openAPI, _ := utils.ConvertAPIDoc(swagger)
	// 3. 分析openAPI JSON数据格式，生成对应的文件
	genFilesConfig := public.GenFilesConfig{
		OutDir:  config.OutDir,
		OpenAPI: openAPI,
	}
	utils.GenerateAPIFiles(genFilesConfig)
}
