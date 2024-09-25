package main

import (
	"example.com/m/v2/src/lib"
	"example.com/m/v2/src/public"
	"flag"
	"fmt"
	"os"
)

func main() {
	// 声明命令行参数
	docUrl := flag.String("docUrl", "", "openAPI 在线文档地址")
	requestLibPath := flag.String("requestLibPath", "", "请求库的引用路径(例如：import request from '@/utils/http;' ")
	outDir := flag.String("outDir", "", "API文档的输出路径，默认是在根目录下(非必传)")

	// 解析命令行参数
	flag.Parse()

	// 检查必传参数
	if *docUrl == "" || *requestLibPath == "" {
		fmt.Println("docUrl, requestLibPath参数必传，请检查！")
		flag.Usage() // 输出参数使用说明
		os.Exit(1)   // 退出程序
	}

	//docUrl := "https://promonkeyli.top:8080/swagger/doc.json"
	//docUrl := "https://petstore.swagger.io/v2/swagger.json"
	config := public.ApiConfig{
		DocUrl:         *docUrl,
		RequestLibPath: *requestLibPath,
		OutDir:         *outDir,
	}
	lib.GenerateAPI(config)
}
