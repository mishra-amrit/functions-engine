package handler

import (
	"encoding/json"
	. "myappengine.io/util"
	"net/http"
	"myappengine.io/ds"
)

var path = "/heartbeat"

type HeartbeatResponse struct {
	Message string `json:"message"`
}

/*
	The attach function which attaches the handler to the server.
	It is always exported in the Go file, for each handler.
	It is the only function in a handler which is exported.
*/
func AttachHeartbeatHandler(appEngineConfig ds.AppEngineConfig) {
	http.HandleFunc(appEngineConfig.ContextPath+path, heartbeatHandler)
}

/*
	The handler is not exported, it's attached to the server by AttachHeartbeatHandler()
*/
func heartbeatHandler(respWriter http.ResponseWriter, request *http.Request) {

	Log("Received : " + request.Method + path)

	if request.Method=="GET" {
		heartbeatResp := HeartbeatResponse{Message: "Service is up !"}
		response, err := json.Marshal(heartbeatResp)

		if err != nil {
			http.Error(respWriter, err.Error(), 500)
			return
		}

		respWriter.Header().Set("content-type", "application/json")
		respWriter.Write(response)
		return
	}else{
		respWriter.WriteHeader(405)
	}

}
