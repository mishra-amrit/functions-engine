package boot

import (
	"os"
	"fmt"
	"encoding/json"
)

func LoadAppEngineConfiguration(file string) AppEngineConfig {
	var config AppEngineConfig
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}