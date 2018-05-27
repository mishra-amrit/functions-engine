/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package app

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"myappengine.io/ds"
	. "myappengine.io/util"
)

func LoadConfiguration(file string) ds.AppEngineConfig {

	var config ds.AppEngineConfig

	data, err := ioutil.ReadFile(file)
	Log(string(data[:]))
	if err != nil {
		Log(err.Error())
	}
	yaml.Unmarshal(data, &config)

	return config

}

func Initialize() ds.AppEngineEnv {

	Log("Loading Configurations..")

	appEngineEnv := ds.AppEngineEnv{}
	appEngineConfig := LoadConfiguration("configuration/myappengine_config.yaml")

	Log("LogPath \t: " + appEngineConfig.LogPath)
	Log("DeploymentPath \t: " + appEngineConfig.DeploymentPath)

	appEngineEnv.AppEngineConfig = appEngineConfig

	Log("Configurations Loaded.")

	return appEngineEnv

}
