# myAppEngine runtime for Java 8
Performs the execution of Java apps uploaded to myAppEngine. This forms as a base for the Java apps.

The target class of App should have one or more of the following method signatures :
	
    public <ResponseDTO-Class> <method-name>()
    
    public <ResponseDTO-Class> <method-name>(HashMap<String, Object> queryParams)
    
    public <ResponseDTO-Class> <method-name>(HashMap<String, Object> queryParams, 
						HashMap<String, Object> pathParams)
    
    public <ResponseDTO-Class> <method-name>(HashMap<String, Object> queryParams, 
						HashMap<String, Object> pathParams, 
						<RequestDTO-Class> payload)

The below requests will be pushed from MyAppEngine to the runtime for registration and execution of App.
                          		
To register a jar :

    POST http://127.0.0.1:8010/myappengine-runtime-java/v1/load
    Body : 
    {
      "appName":"TestApp",
      "jarFilePath":"/Users/amritmishra/Documents/workspace-open/TestProject/build/libs/TestProject.jar",
      "className":"com.aklabs.TestPrg",
      "methodName":"process",
      "requestPayloadClassName":"com.aklabs.RequestDTO",
      "responsePayloadClassName":"com.aklabs.ResponseDTO"
    }
	
To re-register a jar :

    PUT http://127.0.0.1:8010/myappengine-runtime-java/v1/load
    Body : 
    {
      "appName":"TestApp",
      "jarFilePath":"/Users/amritmishra/Documents/workspace-open/TestProject/build/libs/TestProject.jar",
      "className":"com.aklabs.TestPrg",
      "methodName":"process",
      "requestPayloadClassName":"com.aklabs.RequestDTO",
      "responsePayloadClassName":"com.aklabs.ResponseDTO"
    }

To execute the method :

    POST http://127.0.0.1:8010/myappengine-runtime-java/v1/exec/TestApp
    {
      "queryParams":[{"key":"param1","value":"value1"},{"key":"param2","value":"value2"}],
      "pathParams":[{"key":"param3","value":"value3"},{"key":"param4","value":"value4"}],
      "payload":{
        "a":4,
        "b":5
      }
    }
    
    For requests without query or path params :
    {
	  "queryParams":[],
	  "pathParams":[],
	  "payload":{
	    "a":999,
	    "b":1
	  }
	}

Method signature of TestApp :
	
    public ResponseDTO process()
    
    public ResponseDTO process(HashMap<String, Object> queryParams)
    
    public ResponseDTO process(HashMap<String, Object> queryParams, 
                               HashMap<String, Object> pathParams)
    
    public ResponseDTO process(HashMap<String, Object> queryParams, 
                               HashMap<String, Object> pathParams, 
                               RequestDTO payload)

RequestDTO.java

	package com.aklabs;
	
	public class RequestDTO {
	
	    private Integer a;
	    private Integer b;
	
	    public Integer getA() {
		return a;
	    }
	
	    public void setA(Integer a) {
		this.a = a;
	    }
	
	    public Integer getB() {
		return b;
	    }
	
	    public void setB(Integer b) {
		this.b = b;
	    }
	
	    @Override
	    public String toString() {
		return "RequestDTO [a=" + a + ", b=" + b + "]";
	    }
	
	}
	
ResponseDTO.java

	package com.aklabs;
	
	public class ResponseDTO {
	
	    private Integer result;
	
	    public ResponseDTO() {
	    }
	
	    public ResponseDTO(Integer result) {
		this.result = result;
	    }
	
	    public Integer getResult() {
		return result;
	    }
	
	    public void setResult(Integer result) {
		this.result = result;
	    }
	
	    @Override
	    public String toString() {
		return "ResponseDTO [result=" + result + "]";
	    }
	
	}


TestPrg.java

	package com.aklabs;
	
	import java.util.HashMap;
	
	public class TestPrg {
	
	    public ResponseDTO process() {
		return new ResponseDTO(0);
	    }
	
	    public ResponseDTO process(HashMap<String, Object> queryParams) {
		return new ResponseDTO(1);
	    }
	
	    public ResponseDTO process(HashMap<String, Object> queryParams, HashMap<String, Object> pathParams) {
		return new ResponseDTO(2);
	    }
	
	    public ResponseDTO process(HashMap<String, Object> queryParams, HashMap<String, Object> pathParams,
	            RequestDTO payload) {
		Integer result = payload.getA() + payload.getB();
		return new ResponseDTO(result);
	    }
	
	}
    
    
    
	
