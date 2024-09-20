package utils

import (
	"example.com/m/v2/src/public"
	"github.com/getkin/kin-openapi/openapi3"
	"strings"
)

type ApiTypeData struct {
	Name          string
	Type          string
	Description   string
	Properties    []Property
	Enum          []any
	EnumLastIndex int
}
type Property struct {
	Name        string
	Description string
	Type        string
}

func CreateApiType(o openapi3.T, outDir string) {
	// 1. 提取 schemas, swagger2中称为 definitions, openAPI3 更换为 components => schemas
	schemas := o.Components.Schemas

	var TplData []ApiTypeData
	for key, schemaRef := range schemas {
		var name string
		if strings.Contains(key, ".") {
			name = strings.Split(key, ".")[1]
		} else {
			name = key
		}
		schema := *schemaRef.Value
		typeArr := *schema.Type
		enumArr := schema.Enum          // 枚举
		properties := schema.Properties // 该类型的一些属性

		var propertyArr []Property
		for pName, pSchemaRef := range properties {
			if pSchemaRef.Value == nil {
				// 这里可能复用已有类型
				refLink := pSchemaRef.Ref
				if refLink != "" {
					// 解析引用对应的对象

				}
			} else {

			}
			pSchema := *pSchemaRef.Value
			var pTypeArr []string
			if pSchema.Type == nil {
				// 字段没有类型，又可能是 扩展类型，需要从allOf中取类型名称
				if len(pSchema.AllOf) > 0 {
					// allOf ref 示例：#/components/schemas/utils.StatusCode
					ref := pSchema.AllOf[0].Ref
					refTypeStrings := strings.Split(ref, "/")
					refTypeString := refTypeStrings[len(refTypeStrings)-1]
					// 此处需要判断string中是否包含. 包含的话 要分割后取最后一个作为类型名称
					var t string
					if strings.Contains(refTypeString, ".") {
						t = strings.Split(refTypeString, ".")[1]
					} else {
						t = refTypeString
					}
					pTypeArr = append(pTypeArr, t)
				} else {
					pTypeArr = []string{}
				}
			} else {
				pTypeArr = *pSchema.Type
			}

			var pType string
			if len(pTypeArr) > 0 {
				pType = pTypeArr[0]
			} else {
				pType = ""
			}
			propertyArrItem := Property{
				Name:        pName,
				Description: pSchema.Description,
				// 这里的类型需要做一下转换，转成js语言中的类型
				Type: handlerType(pType),
			}
			propertyArr = append(propertyArr, propertyArrItem)
		}

		TplDataItem := ApiTypeData{
			Name:          name,
			Description:   schema.Description,
			Type:          handlerType(typeArr[0]),
			Properties:    propertyArr,
			Enum:          enumArr,
			EnumLastIndex: len(enumArr) - 1,
		}
		TplData = append(TplData, TplDataItem)
	}
	// 模版配置数据组装
	api := public.TplConfig{
		TplPath:     "src/template/api_type.txt",
		TplName:     "api_type.txt",
		TplData:     TplData,
		OutFileName: outDir + "/typings.d.ts",
	}
	RenderTpl(api)
}

func handlerType(t string) string {
	if t == "integer" {
		return "number"
	}
	if t == "" {
		return "any"
	}
	return t
}
