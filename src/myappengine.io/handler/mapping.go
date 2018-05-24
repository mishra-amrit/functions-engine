package handler

import (
	. "myappengine.io/ds"
	. "myappengine.io/util"
)

var handlers []HandlerConfig

func LoadMappings() {

	Log("Loading Mappings..")
	handlers = append(handlers, HandlerConfig{GET, "/app/load"})
	Log("Mappings Loaded.")

}
