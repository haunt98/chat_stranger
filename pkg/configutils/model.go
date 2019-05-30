package configutils

type Environment struct {
	EnvName     string `json:"Name"`
	ServiceName string
	ProjectName string
	ConfigJson  string
}

type ConfigServiceResponse struct {
	ReturnCode    int         `json:"returncode"`
	ReturnMessage string      `json:"returnmessage"`
	Data          Environment `json:"data"`
}
