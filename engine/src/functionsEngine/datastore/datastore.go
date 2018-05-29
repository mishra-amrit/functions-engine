package datastore

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"functionsEngine/structs"
	. "functionsEngine/util"
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
			loadAppRepository(dataPath)
			isDsInitialized = true
		} else {
			Log("File does not exist")
			appRepoFileBuffer, marshalErr := yaml.Marshal(appRepo)
			if marshalErr == nil {
				ioutil.WriteFile(config.DataPath + "/" + appRepoFileName, appRepoFileBuffer, 0666)
			}
		}
	} else {
		Log("DataPath is not set")
	}
}

func loadAppRepository(appRepoLocation string) bool {
	appRepoFileBuffer, err := ioutil.ReadFile(appRepoLocation + "/" + appRepoFileName)
	if err == nil {
		yamlReadErr := yaml.Unmarshal(appRepoFileBuffer, appRepo)
		if yamlReadErr == nil {
			return true
		}
	}

	return false
}

func GetAppRepository() AppRepo {
	return appRepo
}

func CommitAppRepository(){
	mu.Lock()
	appRepoFileBuffer, marshalErr := yaml.Marshal(appRepo)
	if marshalErr == nil {
		ioutil.WriteFile(dataPath + "/" + appRepoFileName, appRepoFileBuffer, 0666)
	}
	mu.Unlock()
}
