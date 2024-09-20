package main

import (
	"example.com/m/v2/src/lib"
	"example.com/m/v2/src/public"
)

func main() {
	//docUrl := "https://promonkeyli.top:8080/swagger/doc.json"
	docUrl := "https://petstore.swagger.io/v2/swagger.json"
	config := public.ApiConfig{
		DocUrl:         docUrl,
		RequestLibPath: "",
		OutDir:         "",
	}
	lib.GenerateAPI(config)
}
