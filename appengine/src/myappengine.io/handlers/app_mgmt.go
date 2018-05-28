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

package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"myappengine.io/structs"
	. "myappengine.io/util"
	"net/http"
)

/*  App Create / Update Request  */
type AppMgmtRequest struct {
	AppName                  string `json:"appName"`
	AppLang                  string `json:"appLang"`
	ClassName                string `json:"className"`
	MethodName               string `json:"methodName"`
	RequestPayloadClassName  string `json:"requestPayloadClassName"`
	ResponsePayloadClassName string `json:"responsePayloadClassName"`
}

/*  App Create / Update Resposne  */
type AppMgmtResponse struct {
	Message string `json:"message"`
}

/*
	The attach function which attaches the handlers to the server.
	It is always exported in the Go file, for each handlers.
	It is the only function in a handlers which is exported.
*/
func AttachAppRegistrationHandler(appEngineConfig structs.AppEngineConfig, router *mux.Router) {
	createAppPath := "/app"
	(*router).HandleFunc(appEngineConfig.ContextPath+createAppPath, appCreationHandler).Methods("POST")
	updateAppPath := "/app/{appName}"
	(*router).HandleFunc(appEngineConfig.ContextPath+updateAppPath, appUpdationHandler).Methods("PUT")
}

/*
	The handlers is not exported, it's attached to the server by AttachAppRegistrationHandler()
*/
func appCreationHandler(respWriter http.ResponseWriter, request *http.Request) {

	Log("appCreationHandler \t:: received : " + request.Method + "\t" + request.RequestURI)

	heartbeatResp := AppMgmtResponse{Message: "Service is up !"}
	response, err := json.Marshal(heartbeatResp)

	if err != nil {
		http.Error(respWriter, err.Error(), 500)
		return
	}

	respWriter.Header().Set("content-type", "application/json")
	respWriter.Write(response)

	return

}

/*
	The handlers is not exported, it's attached to the server by AttachAppRegistrationHandler()
*/
func appUpdationHandler(respWriter http.ResponseWriter, request *http.Request) {

	Log("appUpdationHandler \t:: received : " + request.Method + "\t" + request.RequestURI)

	heartbeatResp := AppMgmtResponse{Message: "Service is up !"}
	response, err := json.Marshal(heartbeatResp)

	if err != nil {
		http.Error(respWriter, err.Error(), 500)
		return
	}

	respWriter.Header().Set("content-type", "application/json")
	respWriter.Write(response)
	return

}
