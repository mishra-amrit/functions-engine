package handlers

import (
	"encoding/json"
	"myappengine.io/ds"
	. "myappengine.io/util"
	"net/http"
	"github.com/gorilla/mux"
)

/*
	The attach function which attaches the handlers to the server.
	It is always exported in the Go file, for each handlers.
	It is the only function in a handlers which is exported.
*/
func AttachAppInvokeHandler(appEngineConfig ds.AppEngineConfig, router *mux.Router) {
	path := "/app/{appName}"
	(*router).HandleFunc(appEngineConfig.ContextPath+path, appInvokeHandler)
}

/*
	The handlers is not exported, it's attached to the server by AttachRegisterAppHandler()
*/
func appInvokeHandler(respWriter http.ResponseWriter, request *http.Request) {

	Log("Received : " + request.Method + "\t" + request.RequestURI)

	routeVars := mux.Vars(request)
	Log("appName : "+routeVars["appName"])

	if request.Method == "GET" {
		heartbeatResp := ds.RegisterAppResponse{Message: "Service is up !"}
		response, err := json.Marshal(heartbeatResp)

		if err != nil {
			http.Error(respWriter, err.Error(), 500)
			return
		}

		respWriter.Header().Set("content-type", "application/json")
		respWriter.Write(response)
		return
	} else {
		respWriter.WriteHeader(405)
	}

}
