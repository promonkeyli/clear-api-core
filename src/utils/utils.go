package utils

import (
	"github.com/getkin/kin-openapi/openapi3"
	"os"
)

type ClassByTag struct {
	Tag   string
	Paths openapi3.Paths
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
func dividePathsByTag(o openapi3.T) map[string]ClassByTag {
	ClassByTagObj := make(map[string]ClassByTag)

	var tags []string // 所有的tag
	for _, tag := range o.Tags {
		tags = append(tags, tag.Name)
	}

	return ClassByTagObj
}
