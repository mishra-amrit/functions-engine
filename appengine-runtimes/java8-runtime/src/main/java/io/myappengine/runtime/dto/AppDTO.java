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

package io.myappengine.runtime.dto;

public class AppDTO {

    private String appName;
    private String jarFilePath;
    private String className;
    private String methodName;
    private String requestPayloadClassName;
    private String responsePayloadClassName;

    public String getJarFilePath() {
	return jarFilePath;
    }

    public void setJarFilePath(String jarFilePath) {
	this.jarFilePath = jarFilePath;
    }

    public String getClassName() {
	return className;
    }

    public void setClassName(String className) {
	this.className = className;
    }

    public String getMethodName() {
	return methodName;
    }

    public void setMethodName(String methodName) {
	this.methodName = methodName;
    }

    public String getAppName() {
	return appName;
    }

    public void setAppName(String appName) {
	this.appName = appName;
    }

    public String getRequestPayloadClassName() {
	return requestPayloadClassName;
    }

    public void setRequestPayloadClassName(String payloadClassName) {
	this.requestPayloadClassName = payloadClassName;
    }

    public String getResponsePayloadClassName() {
	return responsePayloadClassName;
    }

    public void setResponsePayloadClassName(String responseClassName) {
	this.responsePayloadClassName = responseClassName;
    }

}
