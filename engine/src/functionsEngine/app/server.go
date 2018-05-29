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
	"github.com/gorilla/mux"
	"functionsEngine/handlers"
	"functionsEngine/structs"
	"net/http"
)

func StartServer(appEngineConfig structs.AppEngineConfig) {

	router := mux.NewRouter()

	/* Attach handlers by invoking the attachHandler functions */
	handlers.AttachAppEngineHeartbeatHandler(appEngineConfig, router)
	handlers.AttachAppRegistrationHandler(appEngineConfig, router)
	handlers.AttachAppListHandler(appEngineConfig, router)
	handlers.AttachAppInvokeHandler(appEngineConfig, router)

	/* Start the listener */
	http.ListenAndServe(":"+appEngineConfig.ServerPort, router)

}
