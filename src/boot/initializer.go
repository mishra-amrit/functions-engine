package boot

import (
	"fmt"
)

func InitializeAppEngine() AppEngineEnv{

	fmt.Println("Loading Configurations..")

	appEngineEnv := AppEngineEnv{}
	appEngineConfig := LoadAppEngineConfiguration("configuration/appengine_config.json")

	fmt.Println("LogPath \t: "+appEngineConfig.LogPath)
	fmt.Println("DeploymentPath \t: "+appEngineConfig.DeploymentPath)

	appEngineEnv.AppEngineConfig = appEngineConfig

	fmt.Println("Configurations Loaded.")

	return appEngineEnv

}