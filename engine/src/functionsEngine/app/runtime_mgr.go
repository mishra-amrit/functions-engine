package app

import (
	"functionsEngine/structs"
	"functionsEngine/util"
	"os"
	"os/exec"
	"syscall"
)

func InitRuntimes(appEngineConfig structs.AppEngineConfig) {
	util.Log("Initializing Runtimes...")
	initJavaRuntime(appEngineConfig)
	util.Log("Runtimes initialized.")
}

var bytes []byte

func initJavaRuntime(appEngineConfig structs.AppEngineConfig) {
	util.Log("Java")
	binary, lookErr := exec.LookPath("java")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"java", "-jar", appEngineConfig.JavaRuntimeExec," &"}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
