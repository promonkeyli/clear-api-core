package utils

import (
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
)

func convertSwagger2ToOpenAPI3(swaggerDoc string) (string, error) {
	// 解析 Swagger 2.0 JSON 文档
	var swagger openapi2.T
	if err := json.Unmarshal([]byte(swaggerDoc), &swagger); err != nil {
		return "", fmt.Errorf("failed to unmarshal Swagger 2.0: %w", err)
	}
	Log("swagger 2.0 JSON 文档解析完成!", Green)
	// 转换为 OpenAPI 3.0
	openapi3Doc, err1 := openapi2conv.ToV3(&swagger)
	if err1 != nil {
		return "", fmt.Errorf("failed to convert to OpenAPI 3.0: %w", err1)
	}
	// 将 OpenAPI 3.0 文档序列化为 JSON
	openapi3JSON, err2 := json.MarshalIndent(openapi3Doc, "", "  ")
	if err2 != nil {
		return "", fmt.Errorf("failed to marshal OpenAPI 3.0: %w", err2)
	}
	Log("openAPI 3.0 JSON 文档转换完成!", Green)
	return string(openapi3JSON), nil
}

func ConvertAPIDoc(jsonDoc string) (string, error) {
	// 检查 JSON 文档的根字段是 "swagger" 还是 "openapi"
	var check map[string]interface{}
	if err := json.Unmarshal([]byte(jsonDoc), &check); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	if _, isSwagger := check["swagger"]; isSwagger {
		// 如果是 Swagger 2.0，则进行转换
		Log("检测到 swagger 2.0, OpenAPI 3.0转换中...", Purple)
		return convertSwagger2ToOpenAPI3(jsonDoc)
	} else if _, isOpenAPI := check["openapi"]; isOpenAPI {
		// 如果是 OpenAPI 3.0，则不进行转换
		Log("检测到 openAPI 3.0, 没有需要进行转换的!", Purple)
		return jsonDoc, nil
	}

	return "", fmt.Errorf("unrecognized API document format")
}
