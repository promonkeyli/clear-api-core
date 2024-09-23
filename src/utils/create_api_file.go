package utils

import (
	"example.com/m/v2/src/public"
	"github.com/getkin/kin-openapi/openapi3"
)

type ApiServiceData struct {
	// 接口描述
	Description string
	// 接口方法名
	FuncName string
}
type ApiFileData struct {
	ControllerName string
	FileName       string
	RequestLibPath string
	ServiceData    []ApiServiceData
}

func CreateApiFiles(o openapi3.T, outDir string, requestLibPath string) {
	var data []ApiFileData

	// 按 tag 分类paths
	dividePathsByTag(o)

	for _, tag := range o.Tags {
		data = append(data, ApiFileData{
			ControllerName: tag.Name,
			FileName:       tag.Name,
			RequestLibPath: requestLibPath,
			ServiceData:    handleApiServiceData(tag, o),
		})
	}
	// 组织数据

	// 多个接口文件，需要循环生成
	for _, item := range data {
		// 模版配置数据组装
		api := public.TplConfig{
			TplPath:     "src/template/api_service.txt",
			TplName:     "api_service.txt",
			TplData:     item,
			OutFileName: outDir + "/" + item.FileName + ".ts",
		}
		RenderTpl(api)
	}

}

// 处理ts service生成-函数
func handleApiServiceData(tag *openapi3.Tag, o openapi3.T) []ApiServiceData {
	return []ApiServiceData{}
}
