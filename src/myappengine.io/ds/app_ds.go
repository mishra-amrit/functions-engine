package ds

type AppData struct {
	AppId       string            `json:"appId"`
	AppName     string            `json:"appName"`
	AppLang     string            `json:"appLang"`
	AppLocation string            `json:"appExecLocation"`
	AppConfig   JavaAppConfigData `json:"javaAppConfig"`
}

type JavaAppConfigData struct {
	AppId                    string   `json:"appId"`
	className                string   `json:"className"`
	methodName               string   `json:"methodName"`
	invokationURL            string   `json:"invokationURL"`
	queryParams              []string `json:"queryParams"`
	pathParams               []string `json:"pathParams"`
	requestPayloadClassName  string   `json:"requestPayloadClassName"`
	responsePayloadClassName string   `json:"responsePayloadClassName"`
}
