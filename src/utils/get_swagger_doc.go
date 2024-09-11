package utils

import (
	"fmt"
	"io"
	"net/http"
)

func GetSwaggerDoc(url string) string {
	res, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println("Error:", err1)
		return ""
	}
	defer func(Body io.ReadCloser) {
		err2 := Body.Close()
		if err2 != nil {
			fmt.Println("Error:", err2)
		}
	}(res.Body)
	// 读取响应体
	body, err3 := io.ReadAll(res.Body)
	if err3 != nil {
		fmt.Println("Error:", err3)
		return ""
	}
	// swagger获取日志记录
	Log("swagger 文档获取完成!", Green)
	return string(body)
}
