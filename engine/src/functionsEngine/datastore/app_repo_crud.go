package datastore

type AppRepo struct {
	Apps map[string]AppConfigData `yaml:"apps"`
}

func (appRepo AppRepo) Add(appConfigData AppConfigData) OperationResult {
	/* Check if the app is already registered. */
	for appName := range appRepo.Apps {
		/* If app is already registered, then skip */
		if appName == appConfigData.AppName {
			return OperationResult{Message: "App is already registered.", Status: false}
		}
	}

	/* If the app is not registered, then register the app */
	appRepo.Apps[appConfigData.AppName] = appConfigData

	return OperationResult{Message: "App registered", Status: true}
}

func (appRepo AppRepo) Delete(targetAppName string) OperationResult {
	/* Check if the app exists */
	for appName := range appRepo.Apps {
		if appName == targetAppName {
			delete(appRepo.Apps, targetAppName)
			return OperationResult{Message: "App deleted.", Status: true}
		}
	}

	return OperationResult{Message: "No such app exists.", Status: false}
}

func (appRepo AppRepo) Get(targetAppName string) *AppConfigData {
	/* Check if the app exists */
	for appName, appConfigData := range appRepo.Apps {
		if appName == targetAppName {
			return &appConfigData
		}
	}

	return nil
}
