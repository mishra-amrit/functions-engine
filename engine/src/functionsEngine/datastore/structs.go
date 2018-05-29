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

package datastore

type AppConfigData struct {
	AppId         string            `yaml:"appId"`
	AppName       string            `yaml:"appName"`
	AppLang       string            `yaml:"appLang"`
	AppLocation   string            `yaml:"appExecLocation"`
	JavaAppConfig JavaAppConfigData `yaml:"javaAppConfig"`
}

type JavaAppConfigData struct {
	AppId                    string   `yaml:"appId"`
	className                string   `yaml:"className"`
	methodName               string   `yaml:"methodName"`
	invokationURL            string   `yaml:"invokationURL"`
	queryParams              []string `yaml:"queryParams"`
	pathParams               []string `yaml:"pathParams"`
	requestPayloadClassName  string   `yaml:"requestPayloadClassName"`
	responsePayloadClassName string   `yaml:"responsePayloadClassName"`
}

type OperationResult struct {
	Message string
	Status  bool
}
