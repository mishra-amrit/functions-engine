package datastore

import (
	"myappengine.io/structs"
	. "myappengine.io/util"
	"os"
)

func InitializeDataStore(config structs.AppEngineConfig) {
	if config.DataPath != "" {
		_, err := os.Stat(config.DataPath + "/" + "app_repo.yaml")
		if !os.IsNotExist(err) {
			Log("File exists")
		} else {
			Log("File does not exist")
		}
	} else {
		Log("DataPath is not set")
	}
}


