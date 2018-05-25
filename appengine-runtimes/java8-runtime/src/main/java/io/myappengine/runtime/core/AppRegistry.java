package io.myappengine.runtime.core;

import java.util.HashMap;

import io.myappengine.runtime.dto.AppDTO;

public class AppRegistry {

    public class AppRegistryData {

	private AppDTO	   appDTO;
	private AppManager appManager;

	public AppDTO getAppDTO() {
	    return appDTO;
	}

	public void setAppDTO(AppDTO appDTO) {
	    this.appDTO = appDTO;
	}

	public AppManager getAppManager() {
	    return appManager;
	}

	public void setAppManager(AppManager appManager) {
	    this.appManager = appManager;
	}

    }

    private static AppRegistry		     instance = null;
    private HashMap<String, AppRegistryData> registry = new HashMap<String, AppRegistryData>();

    private AppRegistry() {
    }

    static {
	instance = new AppRegistry();
    }

    public static AppRegistry getInstance() {
	return instance;
    }

    public void register(AppDTO appDTO, boolean updateIfExists) throws Exception {
	String appName = appDTO.getAppName();
	if (registry.get(appName) != null) {
	    if (!updateIfExists)
		throw new Exception("App with name " + appName + " is already registered.");
	    else {
		AppRegistryData appRegistryData = registry.remove(appName);
		appRegistryData.getAppManager().close();
	    }
	}

	AppManager appManager = new AppManager();
	appManager.bind(appName, appDTO.getJarFilePath(), appDTO.getClassName(), appDTO.getMethodName(),
	        appDTO.getRequestPayloadClassName(), appDTO.getResponsePayloadClassName());

	AppRegistryData appRegistryData = new AppRegistryData();
	appRegistryData.setAppDTO(appDTO);
	appRegistryData.setAppManager(appManager);
	registry.put(appName, appRegistryData);
    }

    public AppManager getAppManager(String appName) {
	AppRegistryData appRegistryData = registry.get(appName);
	if (appRegistryData != null)
	    return appRegistryData.getAppManager();
	else
	    return null;
    }

}
