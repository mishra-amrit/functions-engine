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
	"fmt"
	"functionsEngine/datastore"
	"functionsEngine/structs"
	. "functionsEngine/util"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

/*  App Upload Resposne  */
type AppUploadResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

var deploymentPath string

/*
	The attach function which attaches the handlers to the server.
	It is always exported in the Go file, for each handlers.
	It is the only function in a handlers which is exported.
*/
func AttachUploadAppHandler(appEngineConfig structs.AppEngineConfig, router *mux.Router) {
	updateAppPath := "/app/{appName}/upload"
	deploymentPath = appEngineConfig.DeploymentPath
	(*router).HandleFunc(appEngineConfig.ContextPath+updateAppPath, uploadAppHandler).Methods("POST")
}

/*
	The handler is not exported, it's attached to the server by AttachUploadAppHandler()
*/
func uploadAppHandler(respWriter http.ResponseWriter, request *http.Request) {

	Log("uploadAppHandler \t:: received : " + request.Method + "\t" + request.RequestURI)

	routeVars := mux.Vars(request)
	appName := routeVars["appName"]
	Log("appName : " + appName)

	app := datastore.AppRepository().Get(appName)
	if app != nil {

		request.ParseMultipartForm(32 << 20)
		file, handler, err := request.FormFile("uploadfile")
		if err != nil {
			respond(respWriter, "Error", "")
			return
		}
		fmt.Printf("%v\n", handler.Header)
		defer file.Close()

		/* deploy : The deployment type. Will be "file" for jar  and "dir" for node.js */
		deploy := ""
		extension := ""
		contentType := handler.Header.Get("Content-Type")
		switch contentType {
		case "application/java-archive":
			extension = "jar"
			deploy = "file"
		case "application/zip":
			extension = ""
			deploy = "dir"
		default:
			respond(respWriter, "Error", "The uploaded file type is not allowed")
			return
		}

		finalDeploymentPath := deploymentPath + "/" + appName + "_executable." + extension

		f, err := os.OpenFile(finalDeploymentPath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			respond(respWriter, "Error", "")
			return
		}
		defer f.Close()
		io.Copy(f, file)

		if deploy == "dir" {
			Unzip(finalDeploymentPath, deploymentPath+"/"+appName+"_executable")
		}

		if err != nil {
			http.Error(respWriter, err.Error(), 500)
			respond(respWriter, "Error", "")
			return
		}

		respond(respWriter, "Success", "")
	} else {
		respond(respWriter, "Error", "App is not registered")
	}

	return

}

func respond(respWriter http.ResponseWriter, status string, message string) {
	appUploadResponse := AppUploadResponse{Message: message, Status: status}
	response, _ := json.Marshal(appUploadResponse)
	respWriter.Header().Set("content-type", "application/json")
	respWriter.Write(response)
}
