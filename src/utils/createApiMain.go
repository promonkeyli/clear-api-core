package utils

import "example.com/m/v2/src/public"

type User struct {
	Name  string
	Email string
}

func CreateApiMain() {
	api := public.TplConfig{
		TplPath:     "../tpl/apiMain.txt",
		TplName:     "apiMain.txt",
		TplData:     "",
		OutFileName: "index.ts",
	}
	RenderTpl(api)
}
