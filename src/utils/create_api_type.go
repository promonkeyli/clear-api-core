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
	Required    bool
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
		requiredArr := schema.Required  // 哪些字段必传

		var propertyArr []Property
		for pName, pSchemaRef := range properties {
			var description string
			var pType string

			// 这里可能复用已有类型
			refLink := pSchemaRef.Ref
			if refLink != "" {
				// 解析引用对应的对象，分割后取最后一个值，就是对应的自定义类型
				refTypeStringArr := strings.Split(refLink, "/")
				refTypeString := refTypeStringArr[len(refTypeStringArr)-1]
				pType = refTypeString
			} else {
				pSchema := *pSchemaRef.Value
				description = pSchema.Description
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
				if len(pTypeArr) > 0 {
					// 这里的类型需要做一下转换，转成js语言中的类型
					pType = handlerType(pTypeArr[0], pSchema)
				} else {
					pType = "any"
				}
			}

			// 处理 是否是必须输入的选项
			propertyArrItem := Property{
				Name:        pName,
				Description: description,
				Type:        pType,
				Required:    includeName(pName, requiredArr),
			}
			propertyArr = append(propertyArr, propertyArrItem)
		}

		TplDataItem := ApiTypeData{
			Name:          name,
			Description:   schema.Description,
			Type:          changeType(typeArr[0]),
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
	// openAPI 类型 文件转换记录
	Log("openAPI 类型主文件渲染完成", Green)
}

func handlerType(t string, s openapi3.Schema) string {
	if t == "array" {
		// 如果openAPI的类型是数组的话，那么首先需要看items中ref存不存在，存在则引用该类型值，不存在则取type对应的类型
		var itemsType string
		if s.Items.Ref != "" {
			itemsTypeStringArr := strings.Split(s.Items.Ref, "/")
			itemsType = itemsTypeStringArr[len(itemsTypeStringArr)-1]
		} else {
			vt := *s.Items.Value.Type
			itemsType = vt[0]
		}
		return "Array<" + itemsType + ">"
	}
	if t == "integer" {
		return "number"
	}
	if t == "" {
		return "any"
	}
	return t
}

func changeType(t string) string {
	if t == "integer" {
		return "number"
	}
	if t == "" {
		return "any"
	}
	return t
}

func includeName(target string, arr []string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}
