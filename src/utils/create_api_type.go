package utils

import (
	"example.com/m/v2/src/public"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"strings"
)

type ApiTypeData struct {
	Name        string
	Type        string
	Description string
	Properties  []Property
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
		name := strings.Split(key, ".")[1] // 类型名称
		schema := *schemaRef.Value
		typeArr := *schema.Type
		properties := schema.Properties // 该类型的一些属性

		var propertyArr []Property
		for _, pSchemaRef := range properties {
			pSchema := *pSchemaRef.Value
			var pTypeArr []string
			if *pSchema.Type == nil {
				pTypeArr = []string{}
			} else {
				pTypeArr = *pSchema.Type
			}

			fmt.Println("=====", pTypeArr)

			//var pType string
			//if len(pTypeArr) > 0 {
			//	pType = pTypeArr[0]
			//} else {
			//	pType = ""
			//}
			//propertyArrItem := Property{
			//	Name:        pName,
			//	Description: v.Description,
			//	Type:        t[0],
			//}
			//propertyArr = append(propertyArr,[])
		}

		TplDataItem := ApiTypeData{
			Name:        name,
			Description: schema.Description,
			Type:        typeArr[0],
			Properties:  propertyArr,
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
