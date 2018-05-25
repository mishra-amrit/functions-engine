package ds

type AppEngineConfig struct {
	DeploymentPath string `json:"deploymentPath"`
	LogPath        string `json:"logPath"`
}

type AppEngineEnv struct {
	AppEngineConfig AppEngineConfig
}

type HandlerConfig struct {
	Method        string
	InvokationURL string
}

const (
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
)
