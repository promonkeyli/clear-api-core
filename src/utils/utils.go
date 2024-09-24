package utils

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"os"
	"strings"
)

type OperationByTag struct {
	MethodName string
	Path       string
	Operation  *openapi3.Operation
}

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

// 对openAPI相同的tag的paths进行分类
func dividePathsByTag(o openapi3.T) map[string][]OperationByTag {
	OperationByTagObj := make(map[string][]OperationByTag)

	// 遍历paths
	for path, pathItem := range o.Paths.Map() {
		for methodName, methodObj := range pathItem.Operations() {
			if len(methodObj.Tags) > 0 {
				tag := methodObj.Tags[0]
				OperationByTagObj[tag] = append(OperationByTagObj[tag], OperationByTag{
					MethodName: methodName,
					Path:       path,
					Operation:  methodObj,
				})
			}
		}

	}
	return OperationByTagObj
}

func generateTypeScriptType(schema *openapi3.Schema) string {
	if schema == nil {
		return "unknown" // 处理 nil schema
	}
	sType := *schema.Type

	switch sType[0] {
	case "object":
		tsType := "{ " // 开始对象类型
		for propName, prop := range schema.Properties {
			tsType += fmt.Sprintf("%s: %s; ", propName, generateTypeScriptType(prop.Value))
		}
		tsType += "}"
		return tsType
	case "array":
		return "Array<" + generateTypeScriptType(schema.Items.Value) + ">"
	case "string":
		return "string"
	case "integer":
		return "number"
	case "boolean":
		return "boolean"
	// 添加其他基本类型支持
	default:
		return "unknown" // 处理未知类型
	}
}

func replacePathTpl(path string) string {
	// 替换占位符格式
	return strings.ReplaceAll(path, "{", "${")
}
