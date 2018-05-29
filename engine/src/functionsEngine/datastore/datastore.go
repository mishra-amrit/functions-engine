package datastore

import (
	"functionsEngine/structs"
	. "functionsEngine/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sync"
)

var mu sync.Mutex

const appRepoFileName = "app_repo.yaml"

var dataPath string

var isDsInitialized = false

var appRepo AppRepo

func Initialize(config structs.AppEngineConfig) {
	if config.DataPath != "" {
		dataPath = config.DataPath
		_, err := os.Stat(dataPath + "/" + appRepoFileName)
		if !os.IsNotExist(err) {
			Log("File exists")
			/* Load the appRepository. This contains all the registered app information. */
			if loadAppRepository(dataPath) {
				Log("App Repository Loaded.")
			}
			isDsInitialized = true
		} else {
			Log("File does not exist")
			appRepoFileBuffer, marshalErr := yaml.Marshal(appRepo)
			if marshalErr == nil {
				ioutil.WriteFile(config.DataPath+"/"+appRepoFileName, appRepoFileBuffer, 0666)
			}
		}
	} else {
		Log("DataPath is not set")
	}
}

func loadAppRepository(appRepoLocation string) bool {
	appRepoFileBuffer, fileReadErr := ioutil.ReadFile(appRepoLocation + "/" + appRepoFileName)
	if fileReadErr == nil {
		yamlReadErr := yaml.Unmarshal(appRepoFileBuffer, &appRepo)
		if yamlReadErr == nil {
			return true
		}
	}

	return false
}

func AppRepository() AppRepo {
	return appRepo
}

func Commit() bool {
	mu.Lock()
	appRepoFileBuffer, marshalErr := yaml.Marshal(appRepo)
	if marshalErr == nil {
		commitErr := ioutil.WriteFile(dataPath+"/"+appRepoFileName, appRepoFileBuffer, 0666)
		if commitErr == nil {
			mu.Unlock()
			return true
		}
	}
	mu.Unlock()
	return false
}
