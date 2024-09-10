package public

type ApiConfig struct {
	DocUrl         string
	RequestLibPath string
	OutDir         string
}

type GenFilesConfig struct {
	OpenAPI string
	OutDir  string
}

type TplConfig struct {
	TplName     string
	TplPath     string
	TplData     interface{}
	OutFileName string
}
