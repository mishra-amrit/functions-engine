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
