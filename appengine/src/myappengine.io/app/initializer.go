package app

import (
	"encoding/json"
	. "myappengine.io/ds"
	. "myappengine.io/util"
	"os"
)

func LoadConfiguration(file string) AppEngineConfig {

	var config AppEngineConfig

	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		Log(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config

}

func Initialize() AppEngineEnv {

	Log("Loading Configurations..")

	appEngineEnv := AppEngineEnv{}
	appEngineConfig := LoadConfiguration("configuration/appengine_config.json")

	Log("LogPath \t: " + appEngineConfig.LogPath)
	Log("DeploymentPath \t: " + appEngineConfig.DeploymentPath)

	appEngineEnv.AppEngineConfig = appEngineConfig

	Log("Configurations Loaded.")

	return appEngineEnv

}
