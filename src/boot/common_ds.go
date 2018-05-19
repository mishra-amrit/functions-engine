package boot

type AppEngineConfig struct {
	DeploymentPath 	string `json:"deploymentPath"`
	LogPath 		string `json:"logPath"`
}

type AppEngineEnv struct {
	AppEngineConfig AppEngineConfig
}

