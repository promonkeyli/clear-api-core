package public

type ApiConfig struct {
	DocUrl         string
	OutDir         string
	RequestLibPath string
}

type GenFilesConfig struct {
	OpenAPI        string
	OutDir         string
	RequestLibPath string
}

type TplConfig struct {
	TplName     string
	TplPath     string
	TplData     interface{}
	OutFileName string
}
