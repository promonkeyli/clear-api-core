package utils

import (
	"example.com/m/v2/src/public"
	"github.com/getkin/kin-openapi/openapi3"
)

type ApiMainData struct {
	ControllerName string
	FileName       string
}

func CreateApiMain(o openapi3.T, outDir string) {
	// 处理 openAPI数据, 数组长度不固定，使用切片动态扩容
	var data []ApiMainData
	for _, tag := range o.Tags {
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
		OutFileName: outDir + "/index.ts",
	}
	RenderTpl(api)
	// openAPI index.ts文件转换记录
	Log("openAPI index.ts 主文件渲染完成", Green)
}
