package boot

import (
	"fmt"
)

func InitializeAppEngine() AppEngineEnv{

	fmt.Println("Loading Configurations..")
	appEngineEnv := AppEngineEnv{}
	appEngineConfig := LoadAppEngineConfiguration("configuration/appengine_config.json")
	fmt.Println("LogPath : "+appEngineConfig.LogPath)
	fmt.Println("DeploymentPath : "+appEngineConfig.DeploymentPath)
	appEngineEnv.AppEngineConfig = appEngineConfig
	fmt.Println("Configurations Loaded.")

	return appEngineEnv

}