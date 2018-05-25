package io.myappengine.runtime.core;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.net.MalformedURLException;
import java.net.URL;
import java.net.URLClassLoader;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;

import io.myappengine.runtime.dto.ParamDTO;

public class AppManager extends URLClassLoader {

    boolean	 isAppLoaded		 = false;
    Class<?>	 theClass		 = null;
    Object	 theClassInstance	 = null;
    List<Method> theMethods		 = new ArrayList<Method>();
    Class<?>	 theRequestPayloadClass	 = null;
    Class<?>	 theResponsePayloadClass = null;

    public AppManager() {
	super(new URL[] {});
    }

    private void loadJar(String path) throws MalformedURLException {
	String urlPath = "jar:file://" + path + "!/";
	addURL(new URL(urlPath));
    }

    public void bind(String appName, String appJarPath, String className, String methodName,
            String requestPayloadClassName, String responsePayloadClassName) throws Exception {
	if (!isAppLoaded) {
	    try {
		System.out.println("Load the jar file");
		loadJar(appJarPath);
		System.out.println("Jar file loaded");
		theClass = loadClass(className);

		if (requestPayloadClassName != null)
		    theRequestPayloadClass = loadClass(requestPayloadClassName);

		if (responsePayloadClassName != null)
		    theResponsePayloadClass = loadClass(responsePayloadClassName);

		isAppLoaded = true;
		for (Method method : theClass.getMethods()) {
		    if (method.getName().equals(methodName)) {
			System.out.println(theClass.getName() + "::" + method.getName() + "::paramCount::"
			        + method.getParameterCount());

			for (Class<?> paramType : method.getParameterTypes()) {
			    System.out.println(">> " + paramType.getTypeName());
			}

			theClassInstance = theClass.newInstance();
			theMethods.add(theClass.getMethod(methodName, method.getParameterTypes()));
		    }
		}
		System.out.println("Success!");
	    } catch (Exception ex) {
		System.out.println("Failed.");
		ex.printStackTrace();
	    }
	} else {
	    throw new Exception("AppManager instance is already binded to another app.");
	}

    }

    public Object exec(String appName, List<ParamDTO> pathParamDTOs, List<ParamDTO> queryParamDTOs, Object payload)
            throws IllegalAccessException, IllegalArgumentException, InvocationTargetException {

	if (isAppLoaded) {
	    /* Check the param count */
	    int paramCount = 0;
	    if (pathParamDTOs != null)
		paramCount++;
	    if (queryParamDTOs != null)
		paramCount++;
	    if (payload != null)
		paramCount++;

	    /*
	     * Select the methods from the list of methods which matches the param count
	     */
	    Method theMethod = null;
	    Method defaultMethod = null;
	    for (Method method : theMethods) {
		if (method.getParameterCount() == 0) {
		    defaultMethod = method;
		}
		if (method.getParameterCount() == paramCount) {
		    theMethod = method;
		}
	    }

	    if (theMethod == null && defaultMethod != null)
		theMethod = defaultMethod;
	    else if (theMethod == null && defaultMethod == null)
		return null;

	    /* Prepare the hashmap with param params */
	    Map<String, Object> pathParams = null;
	    if (pathParamDTOs != null) {
		pathParams = new HashMap<String, Object>();
		for (ParamDTO pParamDTO : pathParamDTOs) {
		    pathParams.put(pParamDTO.getKey(), pParamDTO.getValue());
		}
	    }

	    /* Prepare the hashmap with query params */
	    Map<String, Object> queryParams = null;
	    if (queryParamDTOs != null) {
		queryParams = new HashMap<String, Object>();
		for (ParamDTO qParamDTO : queryParamDTOs) {
		    queryParams.put(qParamDTO.getKey(), qParamDTO.getValue());
		}
	    }

	    Object output = null;
	    System.out.println(theClass.getName() + "::" + theMethod.getName() + "::" + " param count :: "
	            + theMethod.getParameterCount());

	    switch (paramCount) {
	    case 0:
		System.out.println(theClass.getName() + "::" + theMethod.getName() + "::" + " without params");
		output = theMethod.invoke(theClassInstance);
		break;
	    case 1:
		System.out.println(theClass.getName() + "::" + theMethod.getName() + "::" + " with query params");
		output = theMethod.invoke(theClassInstance, queryParams);
		break;
	    case 2:
		System.out.println(
		        theClass.getName() + "::" + theMethod.getName() + "::" + " with query and path params");
		output = theMethod.invoke(theClassInstance, queryParams, pathParams);
		break;
	    case 3:
		System.out.println(
		        theClass.getName() + "::" + theMethod.getName() + "::" + " with query, path params & payload");
		if (theRequestPayloadClass != null) {
		    ObjectMapper objectMapper = new ObjectMapper();
		    objectMapper.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
		    output = theMethod.invoke(theClassInstance, queryParams, pathParams,
		            objectMapper.convertValue(payload, theRequestPayloadClass));
		    output = objectMapper.convertValue(output, theResponsePayloadClass);
		} else
		    output = theMethod.invoke(theClassInstance, queryParams, pathParams, payload);

		break;
	    default:
		break;
	    }

	    return output;

	} else
	    return null;

    }

}
