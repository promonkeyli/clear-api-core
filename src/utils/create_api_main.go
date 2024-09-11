package utils

import (
	"encoding/json"
	"example.com/m/v2/src/public"
	"github.com/getkin/kin-openapi/openapi3"
)

type ApiMainData struct {
	ControllerName string
	FileName       string
}

func CreateApiMain(o string) {
	var openAPI *openapi3.T
	// 转json对象
	err := json.Unmarshal([]byte(o), &openAPI)
	if err != nil {
		Log("err"+err.Error(), Red)
		return
	}
	// 处理 openAPI数据, 数组长度不固定，使用切片动态扩容
	var data []ApiMainData
	for _, tag := range openAPI.Tags {
		data = append(data, ApiMainData{
			ControllerName: tag.Name,
			FileName:       tag.Name,
		})
	}
	// 模版配置数据组装
	api := public.TplConfig{
		TplPath:     "src/template/api_main.txt",
		TplName:     "api_main.txt",
		TplData:     data,
		OutFileName: "index.ts",
	}
	RenderTpl(api)
}
