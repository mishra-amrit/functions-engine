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

package main

import (
	"functionsEngine/app"
	"functionsEngine/datastore"
	. "functionsEngine/util"
)

func showBanner() {
	Log("-----------------------------------------------------")
	Log(" FunctionsEngine | version : 0.1 | FaaS | Serverless")
	Log("-----------------------------------------------------")
}

func main() {

	showBanner()

	Log("Initializing Engine..")
	appEngineEnv := app.Initialize()

	Log("Initialize DataStore..")
	datastore.Initialize(appEngineEnv.AppEngineConfig)
	Log("DataStore initialized.")

	go app.InitRuntimes(appEngineEnv.AppEngineConfig)

	Log("Engine Initialized.")

	app.StartServer(appEngineEnv.AppEngineConfig)

}
