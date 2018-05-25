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

package ds

type AppEngineConfig struct {
	DeploymentPath string `yaml:"deploymentPath"`
	DataPath       string `yaml:"dataPath"`
	LogPath        string `yaml:"logPath"`

	ContextPath string `yaml:"contextPath"`
	ServerPort  string `yaml:"serverPort"`
}

type AppEngineEnv struct {
	AppEngineConfig AppEngineConfig
}

type HandlerConfig struct {
	Method        string
	InvokationURL string
}

type RequestMapping struct {
	Method  string
	Path    string
}

const (
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
)
