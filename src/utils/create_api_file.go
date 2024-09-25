package utils

import (
	"example.com/m/v2/src/public"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"strings"
)

type ApiServiceData struct {
	// 接口描述
	Description string
	// 接口方法名
	FuncName string
	// 接口url
	Url string
	// 接口请求方式 post get等
	Method string
	// 请求体是否需要传递 body
	TransferBody bool
	// body 类型
	TransferBodyType string
	// Header Media Type
	MediaType string
	// url path 中是否有params参数
	ParamsInPath bool
	// url中是否有query params参数
	ParamsInQuery bool
	// params 数据、包含 in path
	ParamsInPathArr []Param
	// params 数据、包含 in query
	ParamsInQueryArr []Param
}
type Param struct {
	Name     string
	Type     string
	Required bool
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
	operations := dividePathsByTag(o)

	for _, tag := range o.Tags {
		data = append(data, ApiFileData{
			ControllerName: tag.Name,
			FileName:       tag.Name,
			RequestLibPath: requestLibPath,
			ServiceData:    handleApiServiceData(operations[tag.Name], o),
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
func handleApiServiceData(operations []OperationByTag, o openapi3.T) []ApiServiceData {
	var serviceData []ApiServiceData
	for _, operation := range operations {
		operationObj := operation.Operation

		// 函数名
		var funcName = generateFuncName(operation)

		fmt.Println(funcName)

		// 请求体
		var TransferBody bool
		var TransferBodyType string
		// header mime
		var MediaType string
		if operationObj.RequestBody != nil {
			TransferBody = true
			content := operationObj.RequestBody.Value.Content
			for mediaType, media := range content {
				MediaType = mediaType

				// 解决 body 类型
				if media.Schema.Ref != "" {
					bodyTypeStringArr := strings.Split(media.Schema.Ref, "/")
					bodyTypeString := bodyTypeStringArr[len(bodyTypeStringArr)-1]
					// 手动拼接 API 前缀
					TransferBodyType = "API." + bodyTypeString
				} else {
					if media.Schema != nil {
						TransferBodyType = generateTypeScriptType(media.Schema.Value)
					}
				}

			}
		} else {
			TransferBody = false
		}

		// url 路径处理
		var pathUrl = replacePathTpl(operation.Path)

		// 请求参数 处理
		var ParamsInPath = false
		var ParamsInQuery = false
		var params = operationObj.Parameters
		var ParamsInPathArr []Param
		var ParamsInQueryArr []Param

		if len(params) > 0 {
			for _, param := range params {
				// 如果遍历参数里面有path的参数存在，那就表明需要params参数存在，在前端request库，会自动拼接到url中
				if param.Value.In == "path" {
					ParamsInPath = true
					// param name
					paramName := param.Value.Name
					paramRequired := param.Value.Required
					paramType := generateTypeScriptType(param.Value.Schema.Value)
					ParamsInPathArr = append(ParamsInPathArr, Param{
						Name:     paramName,
						Type:     paramType,
						Required: paramRequired,
					})
				}
				if param.Value.In == "query" {
					ParamsInQuery = true
					// param name
					paramName := param.Value.Name
					paramRequired := param.Value.Required
					paramType := generateTypeScriptType(param.Value.Schema.Value)
					ParamsInQueryArr = append(ParamsInQueryArr, Param{
						Name:     paramName,
						Type:     paramType,
						Required: paramRequired,
					})
				}
			}
		}

		serviceData = append(serviceData, ApiServiceData{
			Description:      operationObj.Description, // 方法描述
			FuncName:         funcName,                 // 客户端函数名称
			Url:              pathUrl,                  // 请求url
			Method:           operation.MethodName,     // 请求方法
			TransferBody:     TransferBody,             // 请求体
			MediaType:        MediaType,                // content-type
			TransferBodyType: TransferBodyType,         // body 类型
			ParamsInPath:     ParamsInPath,             // Url中是否有请求参数
			ParamsInPathArr:  ParamsInPathArr,          // params 参数数组
			ParamsInQuery:    ParamsInQuery,            // url 中是否有查询参数
			ParamsInQueryArr: ParamsInQueryArr,         // url查询参数 数组
		})
	}
	return serviceData
}
