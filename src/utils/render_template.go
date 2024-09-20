package utils

import (
	"example.com/m/v2/src/public"
	"fmt"
	"os"
	"text/template"
)

func RenderTpl(c public.TplConfig) {
	// 1. 解析模版文件
	tpl, err := template.ParseFiles(c.TplPath)
	if err != nil {
		fmt.Println("Error parsing template:", err)
	}
	// 2. 创建输出文件
	outputFile, err2 := os.Create(c.OutFileName)
	if err2 != nil {
		fmt.Println("Error creating output file:", err)
	}
	defer outputFile.Close()
	// 4. 应用模板并将结果写入文件
	err = tpl.ExecuteTemplate(outputFile, c.TplName, c.TplData)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}
