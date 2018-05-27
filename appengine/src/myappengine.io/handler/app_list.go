package handler

import (
	"encoding/json"
	"myappengine.io/ds"
	. "myappengine.io/util"
	"net/http"
)

/*
	The attach function which attaches the handler to the server.
	It is always exported in the Go file, for each handler.
	It is the only function in a handler which is exported.
*/
func AttachAppListHandler(appEngineConfig ds.AppEngineConfig) {
	var path = "/app/list"
	http.HandleFunc(appEngineConfig.ContextPath+path, appListHandler)
}

/*
	The handler is not exported, it's attached to the server by AttachRegisterAppHandler()
*/
func appListHandler(respWriter http.ResponseWriter, request *http.Request) {

	Log("Received : " + request.Method + "\t" + request.RequestURI)

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
